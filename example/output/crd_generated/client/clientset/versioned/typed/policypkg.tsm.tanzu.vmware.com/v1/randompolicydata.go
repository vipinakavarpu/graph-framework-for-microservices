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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	v1 "nexustempmodule/apis/policypkg.tsm.tanzu.vmware.com/v1"
	scheme "nexustempmodule/client/clientset/versioned/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RandomPolicyDatasGetter has a method to return a RandomPolicyDataInterface.
// A group's client should implement this interface.
type RandomPolicyDatasGetter interface {
	RandomPolicyDatas() RandomPolicyDataInterface
}

// RandomPolicyDataInterface has methods to work with RandomPolicyData resources.
type RandomPolicyDataInterface interface {
	Create(ctx context.Context, randomPolicyData *v1.RandomPolicyData, opts metav1.CreateOptions) (*v1.RandomPolicyData, error)
	Update(ctx context.Context, randomPolicyData *v1.RandomPolicyData, opts metav1.UpdateOptions) (*v1.RandomPolicyData, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.RandomPolicyData, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RandomPolicyDataList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RandomPolicyData, err error)
	RandomPolicyDataExpansion
}

// randomPolicyDatas implements RandomPolicyDataInterface
type randomPolicyDatas struct {
	client rest.Interface
}

// newRandomPolicyDatas returns a RandomPolicyDatas
func newRandomPolicyDatas(c *PolicypkgTsmV1Client) *randomPolicyDatas {
	return &randomPolicyDatas{
		client: c.RESTClient(),
	}
}

// Get takes name of the randomPolicyData, and returns the corresponding randomPolicyData object, and an error if there is any.
func (c *randomPolicyDatas) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RandomPolicyData, err error) {
	result = &v1.RandomPolicyData{}
	err = c.client.Get().
		Resource("randompolicydatas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RandomPolicyDatas that match those selectors.
func (c *randomPolicyDatas) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RandomPolicyDataList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RandomPolicyDataList{}
	err = c.client.Get().
		Resource("randompolicydatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested randomPolicyDatas.
func (c *randomPolicyDatas) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("randompolicydatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a randomPolicyData and creates it.  Returns the server's representation of the randomPolicyData, and an error, if there is any.
func (c *randomPolicyDatas) Create(ctx context.Context, randomPolicyData *v1.RandomPolicyData, opts metav1.CreateOptions) (result *v1.RandomPolicyData, err error) {
	result = &v1.RandomPolicyData{}
	err = c.client.Post().
		Resource("randompolicydatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(randomPolicyData).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a randomPolicyData and updates it. Returns the server's representation of the randomPolicyData, and an error, if there is any.
func (c *randomPolicyDatas) Update(ctx context.Context, randomPolicyData *v1.RandomPolicyData, opts metav1.UpdateOptions) (result *v1.RandomPolicyData, err error) {
	result = &v1.RandomPolicyData{}
	err = c.client.Put().
		Resource("randompolicydatas").
		Name(randomPolicyData.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(randomPolicyData).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the randomPolicyData and deletes it. Returns an error if one occurs.
func (c *randomPolicyDatas) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("randompolicydatas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *randomPolicyDatas) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("randompolicydatas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched randomPolicyData.
func (c *randomPolicyDatas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RandomPolicyData, err error) {
	result = &v1.RandomPolicyData{}
	err = c.client.Patch(pt).
		Resource("randompolicydatas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
