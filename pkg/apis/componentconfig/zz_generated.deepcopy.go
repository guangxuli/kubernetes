// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package componentconfig

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupResource) DeepCopyInto(out *GroupResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupResource.
func (in *GroupResource) DeepCopy() *GroupResource {
	if in == nil {
		return nil
	}
	out := new(GroupResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPVar) DeepCopyInto(out *IPVar) {
	*out = *in
	if in.Val != nil {
		in, out := &in.Val, &out.Val
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPVar.
func (in *IPVar) DeepCopy() *IPVar {
	if in == nil {
		return nil
	}
	out := new(IPVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllerManagerConfiguration) DeepCopyInto(out *KubeControllerManagerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ServiceSyncPeriod = in.ServiceSyncPeriod
	out.NodeSyncPeriod = in.NodeSyncPeriod
	out.RouteReconciliationPeriod = in.RouteReconciliationPeriod
	out.ResourceQuotaSyncPeriod = in.ResourceQuotaSyncPeriod
	out.NamespaceSyncPeriod = in.NamespaceSyncPeriod
	out.PVClaimBinderSyncPeriod = in.PVClaimBinderSyncPeriod
	out.MinResyncPeriod = in.MinResyncPeriod
	out.HorizontalPodAutoscalerSyncPeriod = in.HorizontalPodAutoscalerSyncPeriod
	out.HorizontalPodAutoscalerUpscaleForbiddenWindow = in.HorizontalPodAutoscalerUpscaleForbiddenWindow
	out.HorizontalPodAutoscalerDownscaleForbiddenWindow = in.HorizontalPodAutoscalerDownscaleForbiddenWindow
	out.DeploymentControllerSyncPeriod = in.DeploymentControllerSyncPeriod
	out.PodEvictionTimeout = in.PodEvictionTimeout
	out.NodeMonitorGracePeriod = in.NodeMonitorGracePeriod
	out.NodeStartupGracePeriod = in.NodeStartupGracePeriod
	out.NodeMonitorPeriod = in.NodeMonitorPeriod
	out.ClusterSigningDuration = in.ClusterSigningDuration
	out.LeaderElection = in.LeaderElection
	out.VolumeConfiguration = in.VolumeConfiguration
	out.ControllerStartInterval = in.ControllerStartInterval
	if in.GCIgnoredResources != nil {
		in, out := &in.GCIgnoredResources, &out.GCIgnoredResources
		*out = make([]GroupResource, len(*in))
		copy(*out, *in)
	}
	out.ReconcilerSyncLoopPeriod = in.ReconcilerSyncLoopPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllerManagerConfiguration.
func (in *KubeControllerManagerConfiguration) DeepCopy() *KubeControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeControllerManagerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderElectionConfiguration) DeepCopyInto(out *LeaderElectionConfiguration) {
	*out = *in
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderElectionConfiguration.
func (in *LeaderElectionConfiguration) DeepCopy() *LeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(LeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeRecyclerConfiguration) DeepCopyInto(out *PersistentVolumeRecyclerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeRecyclerConfiguration.
func (in *PersistentVolumeRecyclerConfiguration) DeepCopy() *PersistentVolumeRecyclerConfiguration {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeRecyclerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRangeVar) DeepCopyInto(out *PortRangeVar) {
	*out = *in
	if in.Val != nil {
		in, out := &in.Val, &out.Val
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRangeVar.
func (in *PortRangeVar) DeepCopy() *PortRangeVar {
	if in == nil {
		return nil
	}
	out := new(PortRangeVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeConfiguration) DeepCopyInto(out *VolumeConfiguration) {
	*out = *in
	out.PersistentVolumeRecyclerConfiguration = in.PersistentVolumeRecyclerConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeConfiguration.
func (in *VolumeConfiguration) DeepCopy() *VolumeConfiguration {
	if in == nil {
		return nil
	}
	out := new(VolumeConfiguration)
	in.DeepCopyInto(out)
	return out
}
