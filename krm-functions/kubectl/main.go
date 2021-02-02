/*
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"opendev.org/airship/airshipctl/pkg/log"

	"github.com/google/shlex"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	kyaml "sigs.k8s.io/kustomize/kyaml/yaml"
)

const (
	kubectlPath       = "/usr/local/bin/kubectl"
	kubectlContextKey = "--context"
	kubectlContextVar = "${KCTL_CONTEXT}"
	kubectlContext    = kubectlContextKey + "=" + kubectlContextVar
	kubeconfigEnvKey  = "--kubeconfig"
	kubeconfigEnvVar  = "${KUBECONFIG}"
	kubeconfigEnv     = kubeconfigEnvKey + "=" + kubeconfigEnvVar
)

// FnConfig uses for function configuration
type FnConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	KubectlCommand    string `yaml:"kubectlCommand"`
	PrintToStdout     bool   `yaml:"printToStdout,omitempty"`
	MaxRetry          int    `yaml:"maxRetry,omitempty"`
	CheckInterval     int    `yaml:"checkInterval,omitempty"`
}

func main() {
	cfg := &FnConfig{}
	resourceList := &framework.ResourceList{FunctionConfig: &cfg}
	cmd := framework.Command(resourceList, func() error {
		args, err := prepareArgs(cfg)
		if err != nil {
			return err
		}
		log.Print(strings.Join(append(strings.Fields(kubectlPath), args...), " "))
		for i := 0; i < cfg.MaxRetry; i++ {
			clicmd := exec.Command(kubectlPath, args...)
			var out, sterr bytes.Buffer
			clicmd.Stdout = &out
			clicmd.Stderr = &sterr
			err := clicmd.Run()
			if err != nil {
				log.Print(sterr.String())
				log.Print(err)
				log.Printf("Sleep %d seconds and try again\n", cfg.CheckInterval)
				time.Sleep(time.Duration(cfg.CheckInterval) * time.Second)
			} else {
				log.Print(out.String())
				if cfg.PrintToStdout {
					data, err := kyaml.Parse(out.String())
					if err != nil {
						return err
					}
					resourceList.Items = append(resourceList.Items, data)
				}
				break
			}
		}
		return nil
	})
	if err := cmd.Execute(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func prepareArgs(cfg *FnConfig) ([]string, error) {
	args, err := shlex.Split(cfg.KubectlCommand)
	if err != nil {
		return nil, err
	}
	if (os.ExpandEnv(kubectlContextVar) == "") || (os.ExpandEnv(kubeconfigEnvVar) == "") {
		msg := fmt.Sprintf("%v or %v is not set", kubectlContextVar, kubeconfigEnvVar)
		return nil, errors.New(msg)
	}

	if !strings.Contains(cfg.KubectlCommand, kubectlContextKey) {
		args = append(args, os.ExpandEnv(kubectlContext))
	}
	if !strings.Contains(cfg.KubectlCommand, kubeconfigEnvKey) {
		args = append(args, os.ExpandEnv(kubeconfigEnv))
	}

	return args, nil
}
