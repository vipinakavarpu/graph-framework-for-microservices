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

	v1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/crd_generatedapis/gns.tsm.tanzu.vmware.com/v1"
	scheme "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/crd_generatedclient/clientset/versioned/scheme"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AdditionalGnsDatasGetter has a method to return a AdditionalGnsDataInterface.
// A group's client should implement this interface.
type AdditionalGnsDatasGetter interface {
	AdditionalGnsDatas() AdditionalGnsDataInterface
}

// AdditionalGnsDataInterface has methods to work with AdditionalGnsData resources.
type AdditionalGnsDataInterface interface {
	Create(ctx context.Context, additionalGnsData *v1.AdditionalGnsData, opts metav1.CreateOptions) (*v1.AdditionalGnsData, error)
	Update(ctx context.Context, additionalGnsData *v1.AdditionalGnsData, opts metav1.UpdateOptions) (*v1.AdditionalGnsData, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.AdditionalGnsData, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AdditionalGnsDataList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AdditionalGnsData, err error)
	AdditionalGnsDataExpansion
}

// additionalGnsDatas implements AdditionalGnsDataInterface
type additionalGnsDatas struct {
	client rest.Interface
}

// newAdditionalGnsDatas returns a AdditionalGnsDatas
func newAdditionalGnsDatas(c *GnsTsmV1Client) *additionalGnsDatas {
	return &additionalGnsDatas{
		client: c.RESTClient(),
	}
}

// Get takes name of the additionalGnsData, and returns the corresponding additionalGnsData object, and an error if there is any.
func (c *additionalGnsDatas) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.AdditionalGnsData, err error) {
	result = &v1.AdditionalGnsData{}
	err = c.client.Get().
		Resource("additionalgnsdatas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AdditionalGnsDatas that match those selectors.
func (c *additionalGnsDatas) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AdditionalGnsDataList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AdditionalGnsDataList{}
	err = c.client.Get().
		Resource("additionalgnsdatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested additionalGnsDatas.
func (c *additionalGnsDatas) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("additionalgnsdatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a additionalGnsData and creates it.  Returns the server's representation of the additionalGnsData, and an error, if there is any.
func (c *additionalGnsDatas) Create(ctx context.Context, additionalGnsData *v1.AdditionalGnsData, opts metav1.CreateOptions) (result *v1.AdditionalGnsData, err error) {
	result = &v1.AdditionalGnsData{}
	err = c.client.Post().
		Resource("additionalgnsdatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(additionalGnsData).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a additionalGnsData and updates it. Returns the server's representation of the additionalGnsData, and an error, if there is any.
func (c *additionalGnsDatas) Update(ctx context.Context, additionalGnsData *v1.AdditionalGnsData, opts metav1.UpdateOptions) (result *v1.AdditionalGnsData, err error) {
	result = &v1.AdditionalGnsData{}
	err = c.client.Put().
		Resource("additionalgnsdatas").
		Name(additionalGnsData.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(additionalGnsData).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the additionalGnsData and deletes it. Returns an error if one occurs.
func (c *additionalGnsDatas) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("additionalgnsdatas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *additionalGnsDatas) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("additionalgnsdatas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched additionalGnsData.
func (c *additionalGnsDatas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AdditionalGnsData, err error) {
	result = &v1.AdditionalGnsData{}
	err = c.client.Patch(pt).
		Resource("additionalgnsdatas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
