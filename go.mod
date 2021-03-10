module opendev.org/airship/airshipctl

go 1.13

require (
	github.com/Azure/go-autorest/autorest v0.11.7 // indirect
	github.com/Masterminds/sprig/v3 v3.2.0
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/ahmetalpbalkan/dlog v0.0.0-20170105205344-4fb5f8204f26 // indirect
	github.com/ahmetb/dlog v0.0.0-20170105205344-4fb5f8204f26
	github.com/chai2010/gettext-go v0.0.0-20170215093142-bf70f2a70fb1 // indirect
	github.com/containerd/containerd v1.4.1 // indirect
	github.com/docker/docker v1.4.2-0.20200203170920-46ec8731fbce
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/elazarl/goproxy v0.0.0-20190421051319-9d40249d3c2f // indirect
	github.com/elazarl/goproxy/ext v0.0.0-20190421051319-9d40249d3c2f // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/go-git/go-billy/v5 v5.0.0
	github.com/go-git/go-git-fixtures/v4 v4.0.1
	github.com/go-git/go-git/v5 v5.0.0
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/gophercloud/gophercloud v0.6.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20191106031601-ce3c9ade29de // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/lucasjones/reggen v0.0.0-20200904144131-37ba4fa293bb
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	k8s.io/api v0.17.9
	k8s.io/apiextensions-apiserver v0.17.9
	k8s.io/apimachinery v0.17.9
	k8s.io/cli-runtime v0.17.4
	k8s.io/client-go v0.17.9
	k8s.io/kubectl v0.17.4
	opendev.org/airship/go-redfish v0.0.0-20200318103738-db034d1d753a
	opendev.org/airship/go-redfish/client v0.0.0-20200318103738-db034d1d753a
	sigs.k8s.io/cli-utils v0.21.0
	sigs.k8s.io/cluster-api v0.3.13
	sigs.k8s.io/controller-runtime v0.5.14
	sigs.k8s.io/kustomize/api v0.7.2
	sigs.k8s.io/kustomize/kyaml v0.10.6
	sigs.k8s.io/yaml v1.2.0
)

replace k8s.io/kubectl => k8s.io/kubectl v0.0.0-20191219154910-1528d4eea6dd
