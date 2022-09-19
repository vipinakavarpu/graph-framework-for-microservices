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

	v1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/crd_generated/apis/gns.tsm.tanzu.vmware.com/v1"
	scheme "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/crd_generated/client/clientset/versioned/scheme"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GnsesGetter has a method to return a GnsInterface.
// A group's client should implement this interface.
type GnsesGetter interface {
	Gnses() GnsInterface
}

// GnsInterface has methods to work with Gns resources.
type GnsInterface interface {
	Create(ctx context.Context, gns *v1.Gns, opts metav1.CreateOptions) (*v1.Gns, error)
	Update(ctx context.Context, gns *v1.Gns, opts metav1.UpdateOptions) (*v1.Gns, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Gns, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.GnsList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Gns, err error)
	GnsExpansion
}

// gnses implements GnsInterface
type gnses struct {
	client rest.Interface
}

// newGnses returns a Gnses
func newGnses(c *GnsTsmV1Client) *gnses {
	return &gnses{
		client: c.RESTClient(),
	}
}

// Get takes name of the gns, and returns the corresponding gns object, and an error if there is any.
func (c *gnses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Gns, err error) {
	result = &v1.Gns{}
	err = c.client.Get().
		Resource("gnses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Gnses that match those selectors.
func (c *gnses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GnsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.GnsList{}
	err = c.client.Get().
		Resource("gnses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gnses.
func (c *gnses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("gnses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gns and creates it.  Returns the server's representation of the gns, and an error, if there is any.
func (c *gnses) Create(ctx context.Context, gns *v1.Gns, opts metav1.CreateOptions) (result *v1.Gns, err error) {
	result = &v1.Gns{}
	err = c.client.Post().
		Resource("gnses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gns).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gns and updates it. Returns the server's representation of the gns, and an error, if there is any.
func (c *gnses) Update(ctx context.Context, gns *v1.Gns, opts metav1.UpdateOptions) (result *v1.Gns, err error) {
	result = &v1.Gns{}
	err = c.client.Put().
		Resource("gnses").
		Name(gns.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gns).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gns and deletes it. Returns an error if one occurs.
func (c *gnses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("gnses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gnses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("gnses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gns.
func (c *gnses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Gns, err error) {
	result = &v1.Gns{}
	err = c.client.Patch(pt).
		Resource("gnses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
