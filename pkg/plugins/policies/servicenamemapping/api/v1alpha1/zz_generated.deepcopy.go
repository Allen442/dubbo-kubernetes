//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceNameMapping) DeepCopyInto(out *ServiceNameMapping) {
	*out = *in
	if in.ApplicationNames != nil {
		in, out := &in.ApplicationNames, &out.ApplicationNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceNameMapping.
func (in *ServiceNameMapping) DeepCopy() *ServiceNameMapping {
	if in == nil {
		return nil
	}
	out := new(ServiceNameMapping)
	in.DeepCopyInto(out)
	return out
}
