// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Clusterctl) DeepCopyInto(out *Clusterctl) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]*Provider, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Provider)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.InitOptions != nil {
		in, out := &in.InitOptions, &out.InitOptions
		*out = new(InitOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.MoveOptions != nil {
		in, out := &in.MoveOptions, &out.MoveOptions
		*out = new(MoveOptions)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Clusterctl.
func (in *Clusterctl) DeepCopy() *Clusterctl {
	if in == nil {
		return nil
	}
	out := new(Clusterctl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Clusterctl) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitOptions) DeepCopyInto(out *InitOptions) {
	*out = *in
	if in.BootstrapProviders != nil {
		in, out := &in.BootstrapProviders, &out.BootstrapProviders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InfrastructureProviders != nil {
		in, out := &in.InfrastructureProviders, &out.InfrastructureProviders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ControlPlaneProviders != nil {
		in, out := &in.ControlPlaneProviders, &out.ControlPlaneProviders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.KubeConfigRef != nil {
		in, out := &in.KubeConfigRef, &out.KubeConfigRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitOptions.
func (in *InitOptions) DeepCopy() *InitOptions {
	if in == nil {
		return nil
	}
	out := new(InitOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeConfig) DeepCopyInto(out *KubeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeConfig.
func (in *KubeConfig) DeepCopy() *KubeConfig {
	if in == nil {
		return nil
	}
	out := new(KubeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MoveOptions) DeepCopyInto(out *MoveOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MoveOptions.
func (in *MoveOptions) DeepCopy() *MoveOptions {
	if in == nil {
		return nil
	}
	out := new(MoveOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Phase) DeepCopyInto(out *Phase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Phase.
func (in *Phase) DeepCopy() *Phase {
	if in == nil {
		return nil
	}
	out := new(Phase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Phase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhaseConfig) DeepCopyInto(out *PhaseConfig) {
	*out = *in
	if in.ExecutorRef != nil {
		in, out := &in.ExecutorRef, &out.ExecutorRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhaseConfig.
func (in *PhaseConfig) DeepCopy() *PhaseConfig {
	if in == nil {
		return nil
	}
	out := new(PhaseConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhaseGroup) DeepCopyInto(out *PhaseGroup) {
	*out = *in
	if in.Phases != nil {
		in, out := &in.Phases, &out.Phases
		*out = make([]PhaseGroupStep, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhaseGroup.
func (in *PhaseGroup) DeepCopy() *PhaseGroup {
	if in == nil {
		return nil
	}
	out := new(PhaseGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhaseGroupStep) DeepCopyInto(out *PhaseGroupStep) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhaseGroupStep.
func (in *PhaseGroupStep) DeepCopy() *PhaseGroupStep {
	if in == nil {
		return nil
	}
	out := new(PhaseGroupStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhasePlan) DeepCopyInto(out *PhasePlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.PhaseGroups != nil {
		in, out := &in.PhaseGroups, &out.PhaseGroups
		*out = make([]PhaseGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhasePlan.
func (in *PhasePlan) DeepCopy() *PhasePlan {
	if in == nil {
		return nil
	}
	out := new(PhasePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PhasePlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Provider) DeepCopyInto(out *Provider) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Provider.
func (in *Provider) DeepCopy() *Provider {
	if in == nil {
		return nil
	}
	out := new(Provider)
	in.DeepCopyInto(out)
	return out
}
