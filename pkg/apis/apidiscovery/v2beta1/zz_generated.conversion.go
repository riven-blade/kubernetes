//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v2beta1

import (
	unsafe "unsafe"

	v2beta1 "k8s.io/api/apidiscovery/v2beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apidiscovery "k8s.io/kubernetes/pkg/apis/apidiscovery"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v2beta1.APIGroupDiscovery)(nil), (*apidiscovery.APIGroupDiscovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_APIGroupDiscovery_To_apidiscovery_APIGroupDiscovery(a.(*v2beta1.APIGroupDiscovery), b.(*apidiscovery.APIGroupDiscovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apidiscovery.APIGroupDiscovery)(nil), (*v2beta1.APIGroupDiscovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apidiscovery_APIGroupDiscovery_To_v2beta1_APIGroupDiscovery(a.(*apidiscovery.APIGroupDiscovery), b.(*v2beta1.APIGroupDiscovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2beta1.APIGroupDiscoveryList)(nil), (*apidiscovery.APIGroupDiscoveryList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_APIGroupDiscoveryList_To_apidiscovery_APIGroupDiscoveryList(a.(*v2beta1.APIGroupDiscoveryList), b.(*apidiscovery.APIGroupDiscoveryList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apidiscovery.APIGroupDiscoveryList)(nil), (*v2beta1.APIGroupDiscoveryList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apidiscovery_APIGroupDiscoveryList_To_v2beta1_APIGroupDiscoveryList(a.(*apidiscovery.APIGroupDiscoveryList), b.(*v2beta1.APIGroupDiscoveryList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2beta1.APIResourceDiscovery)(nil), (*apidiscovery.APIResourceDiscovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_APIResourceDiscovery_To_apidiscovery_APIResourceDiscovery(a.(*v2beta1.APIResourceDiscovery), b.(*apidiscovery.APIResourceDiscovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apidiscovery.APIResourceDiscovery)(nil), (*v2beta1.APIResourceDiscovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apidiscovery_APIResourceDiscovery_To_v2beta1_APIResourceDiscovery(a.(*apidiscovery.APIResourceDiscovery), b.(*v2beta1.APIResourceDiscovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2beta1.APISubresourceDiscovery)(nil), (*apidiscovery.APISubresourceDiscovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_APISubresourceDiscovery_To_apidiscovery_APISubresourceDiscovery(a.(*v2beta1.APISubresourceDiscovery), b.(*apidiscovery.APISubresourceDiscovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apidiscovery.APISubresourceDiscovery)(nil), (*v2beta1.APISubresourceDiscovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apidiscovery_APISubresourceDiscovery_To_v2beta1_APISubresourceDiscovery(a.(*apidiscovery.APISubresourceDiscovery), b.(*v2beta1.APISubresourceDiscovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2beta1.APIVersionDiscovery)(nil), (*apidiscovery.APIVersionDiscovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2beta1_APIVersionDiscovery_To_apidiscovery_APIVersionDiscovery(a.(*v2beta1.APIVersionDiscovery), b.(*apidiscovery.APIVersionDiscovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apidiscovery.APIVersionDiscovery)(nil), (*v2beta1.APIVersionDiscovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apidiscovery_APIVersionDiscovery_To_v2beta1_APIVersionDiscovery(a.(*apidiscovery.APIVersionDiscovery), b.(*v2beta1.APIVersionDiscovery), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v2beta1_APIGroupDiscovery_To_apidiscovery_APIGroupDiscovery(in *v2beta1.APIGroupDiscovery, out *apidiscovery.APIGroupDiscovery, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Versions = *(*[]apidiscovery.APIVersionDiscovery)(unsafe.Pointer(&in.Versions))
	return nil
}

// Convert_v2beta1_APIGroupDiscovery_To_apidiscovery_APIGroupDiscovery is an autogenerated conversion function.
func Convert_v2beta1_APIGroupDiscovery_To_apidiscovery_APIGroupDiscovery(in *v2beta1.APIGroupDiscovery, out *apidiscovery.APIGroupDiscovery, s conversion.Scope) error {
	return autoConvert_v2beta1_APIGroupDiscovery_To_apidiscovery_APIGroupDiscovery(in, out, s)
}

func autoConvert_apidiscovery_APIGroupDiscovery_To_v2beta1_APIGroupDiscovery(in *apidiscovery.APIGroupDiscovery, out *v2beta1.APIGroupDiscovery, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Versions = *(*[]v2beta1.APIVersionDiscovery)(unsafe.Pointer(&in.Versions))
	return nil
}

// Convert_apidiscovery_APIGroupDiscovery_To_v2beta1_APIGroupDiscovery is an autogenerated conversion function.
func Convert_apidiscovery_APIGroupDiscovery_To_v2beta1_APIGroupDiscovery(in *apidiscovery.APIGroupDiscovery, out *v2beta1.APIGroupDiscovery, s conversion.Scope) error {
	return autoConvert_apidiscovery_APIGroupDiscovery_To_v2beta1_APIGroupDiscovery(in, out, s)
}

func autoConvert_v2beta1_APIGroupDiscoveryList_To_apidiscovery_APIGroupDiscoveryList(in *v2beta1.APIGroupDiscoveryList, out *apidiscovery.APIGroupDiscoveryList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]apidiscovery.APIGroupDiscovery)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v2beta1_APIGroupDiscoveryList_To_apidiscovery_APIGroupDiscoveryList is an autogenerated conversion function.
func Convert_v2beta1_APIGroupDiscoveryList_To_apidiscovery_APIGroupDiscoveryList(in *v2beta1.APIGroupDiscoveryList, out *apidiscovery.APIGroupDiscoveryList, s conversion.Scope) error {
	return autoConvert_v2beta1_APIGroupDiscoveryList_To_apidiscovery_APIGroupDiscoveryList(in, out, s)
}

func autoConvert_apidiscovery_APIGroupDiscoveryList_To_v2beta1_APIGroupDiscoveryList(in *apidiscovery.APIGroupDiscoveryList, out *v2beta1.APIGroupDiscoveryList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v2beta1.APIGroupDiscovery)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_apidiscovery_APIGroupDiscoveryList_To_v2beta1_APIGroupDiscoveryList is an autogenerated conversion function.
func Convert_apidiscovery_APIGroupDiscoveryList_To_v2beta1_APIGroupDiscoveryList(in *apidiscovery.APIGroupDiscoveryList, out *v2beta1.APIGroupDiscoveryList, s conversion.Scope) error {
	return autoConvert_apidiscovery_APIGroupDiscoveryList_To_v2beta1_APIGroupDiscoveryList(in, out, s)
}

func autoConvert_v2beta1_APIResourceDiscovery_To_apidiscovery_APIResourceDiscovery(in *v2beta1.APIResourceDiscovery, out *apidiscovery.APIResourceDiscovery, s conversion.Scope) error {
	out.Resource = in.Resource
	out.ResponseKind = (*v1.GroupVersionKind)(unsafe.Pointer(in.ResponseKind))
	out.Scope = apidiscovery.ResourceScope(in.Scope)
	out.SingularResource = in.SingularResource
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.ShortNames = *(*[]string)(unsafe.Pointer(&in.ShortNames))
	out.Categories = *(*[]string)(unsafe.Pointer(&in.Categories))
	out.Subresources = *(*[]apidiscovery.APISubresourceDiscovery)(unsafe.Pointer(&in.Subresources))
	return nil
}

// Convert_v2beta1_APIResourceDiscovery_To_apidiscovery_APIResourceDiscovery is an autogenerated conversion function.
func Convert_v2beta1_APIResourceDiscovery_To_apidiscovery_APIResourceDiscovery(in *v2beta1.APIResourceDiscovery, out *apidiscovery.APIResourceDiscovery, s conversion.Scope) error {
	return autoConvert_v2beta1_APIResourceDiscovery_To_apidiscovery_APIResourceDiscovery(in, out, s)
}

func autoConvert_apidiscovery_APIResourceDiscovery_To_v2beta1_APIResourceDiscovery(in *apidiscovery.APIResourceDiscovery, out *v2beta1.APIResourceDiscovery, s conversion.Scope) error {
	out.Resource = in.Resource
	out.ResponseKind = (*v1.GroupVersionKind)(unsafe.Pointer(in.ResponseKind))
	out.Scope = v2beta1.ResourceScope(in.Scope)
	out.SingularResource = in.SingularResource
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.ShortNames = *(*[]string)(unsafe.Pointer(&in.ShortNames))
	out.Categories = *(*[]string)(unsafe.Pointer(&in.Categories))
	out.Subresources = *(*[]v2beta1.APISubresourceDiscovery)(unsafe.Pointer(&in.Subresources))
	return nil
}

// Convert_apidiscovery_APIResourceDiscovery_To_v2beta1_APIResourceDiscovery is an autogenerated conversion function.
func Convert_apidiscovery_APIResourceDiscovery_To_v2beta1_APIResourceDiscovery(in *apidiscovery.APIResourceDiscovery, out *v2beta1.APIResourceDiscovery, s conversion.Scope) error {
	return autoConvert_apidiscovery_APIResourceDiscovery_To_v2beta1_APIResourceDiscovery(in, out, s)
}

func autoConvert_v2beta1_APISubresourceDiscovery_To_apidiscovery_APISubresourceDiscovery(in *v2beta1.APISubresourceDiscovery, out *apidiscovery.APISubresourceDiscovery, s conversion.Scope) error {
	out.Subresource = in.Subresource
	out.ResponseKind = (*v1.GroupVersionKind)(unsafe.Pointer(in.ResponseKind))
	out.AcceptedTypes = *(*[]v1.GroupVersionKind)(unsafe.Pointer(&in.AcceptedTypes))
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	return nil
}

// Convert_v2beta1_APISubresourceDiscovery_To_apidiscovery_APISubresourceDiscovery is an autogenerated conversion function.
func Convert_v2beta1_APISubresourceDiscovery_To_apidiscovery_APISubresourceDiscovery(in *v2beta1.APISubresourceDiscovery, out *apidiscovery.APISubresourceDiscovery, s conversion.Scope) error {
	return autoConvert_v2beta1_APISubresourceDiscovery_To_apidiscovery_APISubresourceDiscovery(in, out, s)
}

func autoConvert_apidiscovery_APISubresourceDiscovery_To_v2beta1_APISubresourceDiscovery(in *apidiscovery.APISubresourceDiscovery, out *v2beta1.APISubresourceDiscovery, s conversion.Scope) error {
	out.Subresource = in.Subresource
	out.ResponseKind = (*v1.GroupVersionKind)(unsafe.Pointer(in.ResponseKind))
	out.AcceptedTypes = *(*[]v1.GroupVersionKind)(unsafe.Pointer(&in.AcceptedTypes))
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	return nil
}

// Convert_apidiscovery_APISubresourceDiscovery_To_v2beta1_APISubresourceDiscovery is an autogenerated conversion function.
func Convert_apidiscovery_APISubresourceDiscovery_To_v2beta1_APISubresourceDiscovery(in *apidiscovery.APISubresourceDiscovery, out *v2beta1.APISubresourceDiscovery, s conversion.Scope) error {
	return autoConvert_apidiscovery_APISubresourceDiscovery_To_v2beta1_APISubresourceDiscovery(in, out, s)
}

func autoConvert_v2beta1_APIVersionDiscovery_To_apidiscovery_APIVersionDiscovery(in *v2beta1.APIVersionDiscovery, out *apidiscovery.APIVersionDiscovery, s conversion.Scope) error {
	out.Version = in.Version
	out.Resources = *(*[]apidiscovery.APIResourceDiscovery)(unsafe.Pointer(&in.Resources))
	out.Freshness = apidiscovery.DiscoveryFreshness(in.Freshness)
	return nil
}

// Convert_v2beta1_APIVersionDiscovery_To_apidiscovery_APIVersionDiscovery is an autogenerated conversion function.
func Convert_v2beta1_APIVersionDiscovery_To_apidiscovery_APIVersionDiscovery(in *v2beta1.APIVersionDiscovery, out *apidiscovery.APIVersionDiscovery, s conversion.Scope) error {
	return autoConvert_v2beta1_APIVersionDiscovery_To_apidiscovery_APIVersionDiscovery(in, out, s)
}

func autoConvert_apidiscovery_APIVersionDiscovery_To_v2beta1_APIVersionDiscovery(in *apidiscovery.APIVersionDiscovery, out *v2beta1.APIVersionDiscovery, s conversion.Scope) error {
	out.Version = in.Version
	out.Resources = *(*[]v2beta1.APIResourceDiscovery)(unsafe.Pointer(&in.Resources))
	out.Freshness = v2beta1.DiscoveryFreshness(in.Freshness)
	return nil
}

// Convert_apidiscovery_APIVersionDiscovery_To_v2beta1_APIVersionDiscovery is an autogenerated conversion function.
func Convert_apidiscovery_APIVersionDiscovery_To_v2beta1_APIVersionDiscovery(in *apidiscovery.APIVersionDiscovery, out *v2beta1.APIVersionDiscovery, s conversion.Scope) error {
	return autoConvert_apidiscovery_APIVersionDiscovery_To_v2beta1_APIVersionDiscovery(in, out, s)
}