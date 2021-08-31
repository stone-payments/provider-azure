//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultSecret) DeepCopyInto(out *KeyVaultSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultSecret.
func (in *KeyVaultSecret) DeepCopy() *KeyVaultSecret {
	if in == nil {
		return nil
	}
	out := new(KeyVaultSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyVaultSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultSecretAttributes) DeepCopyInto(out *KeyVaultSecretAttributes) {
	*out = *in
	if in.RecoverableDays != nil {
		in, out := &in.RecoverableDays, &out.RecoverableDays
		*out = new(int32)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.NotBefore != nil {
		in, out := &in.NotBefore, &out.NotBefore
		*out = (*in).DeepCopy()
	}
	if in.Expires != nil {
		in, out := &in.Expires, &out.Expires
		*out = (*in).DeepCopy()
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = (*in).DeepCopy()
	}
	if in.Updated != nil {
		in, out := &in.Updated, &out.Updated
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultSecretAttributes.
func (in *KeyVaultSecretAttributes) DeepCopy() *KeyVaultSecretAttributes {
	if in == nil {
		return nil
	}
	out := new(KeyVaultSecretAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultSecretList) DeepCopyInto(out *KeyVaultSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeyVaultSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultSecretList.
func (in *KeyVaultSecretList) DeepCopy() *KeyVaultSecretList {
	if in == nil {
		return nil
	}
	out := new(KeyVaultSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyVaultSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultSecretObservation) DeepCopyInto(out *KeyVaultSecretObservation) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = new(KeyVaultSecretAttributes)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultSecretObservation.
func (in *KeyVaultSecretObservation) DeepCopy() *KeyVaultSecretObservation {
	if in == nil {
		return nil
	}
	out := new(KeyVaultSecretObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultSecretParameters) DeepCopyInto(out *KeyVaultSecretParameters) {
	*out = *in
	if in.ContentType != nil {
		in, out := &in.ContentType, &out.ContentType
		*out = new(string)
		**out = **in
	}
	if in.SecretAttributes != nil {
		in, out := &in.SecretAttributes, &out.SecretAttributes
		*out = new(KeyVaultSecretAttributes)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultSecretParameters.
func (in *KeyVaultSecretParameters) DeepCopy() *KeyVaultSecretParameters {
	if in == nil {
		return nil
	}
	out := new(KeyVaultSecretParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultSecretSpec) DeepCopyInto(out *KeyVaultSecretSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultSecretSpec.
func (in *KeyVaultSecretSpec) DeepCopy() *KeyVaultSecretSpec {
	if in == nil {
		return nil
	}
	out := new(KeyVaultSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultSecretStatus) DeepCopyInto(out *KeyVaultSecretStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultSecretStatus.
func (in *KeyVaultSecretStatus) DeepCopy() *KeyVaultSecretStatus {
	if in == nil {
		return nil
	}
	out := new(KeyVaultSecretStatus)
	in.DeepCopyInto(out)
	return out
}
