// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentconfig_v1alpha1 "k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&KubeSchedulerConfiguration{}, func(obj interface{}) { SetObjectDefaults_KubeSchedulerConfiguration(obj.(*KubeSchedulerConfiguration)) })
	return nil
}

func SetObjectDefaults_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration) {
	SetDefaults_KubeSchedulerConfiguration(in)
	componentconfig_v1alpha1.SetDefaults_LeaderElectionConfiguration(&in.LeaderElection.LeaderElectionConfiguration)
}
