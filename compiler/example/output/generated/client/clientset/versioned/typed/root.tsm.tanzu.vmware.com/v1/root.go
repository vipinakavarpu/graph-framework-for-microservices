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
	"time"

	v1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/crd_generatedapis/root.tsm.tanzu.vmware.com/v1"
	scheme "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/crd_generatedclient/clientset/versioned/scheme"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RootsGetter has a method to return a RootInterface.
// A group's client should implement this interface.
type RootsGetter interface {
	Roots() RootInterface
}

// RootInterface has methods to work with Root resources.
type RootInterface interface {
	Create(ctx context.Context, root *v1.Root, opts metav1.CreateOptions) (*v1.Root, error)
	Update(ctx context.Context, root *v1.Root, opts metav1.UpdateOptions) (*v1.Root, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Root, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RootList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Root, err error)
	RootExpansion
}

// roots implements RootInterface
type roots struct {
	client rest.Interface
}

// newRoots returns a Roots
func newRoots(c *RootTsmV1Client) *roots {
	return &roots{
		client: c.RESTClient(),
	}
}

// Get takes name of the root, and returns the corresponding root object, and an error if there is any.
func (c *roots) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Root, err error) {
	result = &v1.Root{}
	err = c.client.Get().
		Resource("roots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Roots that match those selectors.
func (c *roots) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RootList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RootList{}
	err = c.client.Get().
		Resource("roots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested roots.
func (c *roots) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("roots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a root and creates it.  Returns the server's representation of the root, and an error, if there is any.
func (c *roots) Create(ctx context.Context, root *v1.Root, opts metav1.CreateOptions) (result *v1.Root, err error) {
	result = &v1.Root{}
	err = c.client.Post().
		Resource("roots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(root).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a root and updates it. Returns the server's representation of the root, and an error, if there is any.
func (c *roots) Update(ctx context.Context, root *v1.Root, opts metav1.UpdateOptions) (result *v1.Root, err error) {
	result = &v1.Root{}
	err = c.client.Put().
		Resource("roots").
		Name(root.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(root).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the root and deletes it. Returns an error if one occurs.
func (c *roots) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("roots").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *roots) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("roots").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched root.
func (c *roots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Root, err error) {
	result = &v1.Root{}
	err = c.client.Patch(pt).
		Resource("roots").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
