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

package versioned

import (
	"fmt"
	configtsmv1 "nexustempmodule/client/clientset/versioned/typed/config.tsm.tanzu.vmware.com/v1"
	gnstsmv1 "nexustempmodule/client/clientset/versioned/typed/gns.tsm.tanzu.vmware.com/v1"
	policypkgtsmv1 "nexustempmodule/client/clientset/versioned/typed/policypkg.tsm.tanzu.vmware.com/v1"
	roottsmv1 "nexustempmodule/client/clientset/versioned/typed/root.tsm.tanzu.vmware.com/v1"
	servicegrouptsmv1 "nexustempmodule/client/clientset/versioned/typed/servicegroup.tsm.tanzu.vmware.com/v1"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ConfigTsmV1() configtsmv1.ConfigTsmV1Interface
	GnsTsmV1() gnstsmv1.GnsTsmV1Interface
	PolicypkgTsmV1() policypkgtsmv1.PolicypkgTsmV1Interface
	RootTsmV1() roottsmv1.RootTsmV1Interface
	ServicegroupTsmV1() servicegrouptsmv1.ServicegroupTsmV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	configTsmV1       *configtsmv1.ConfigTsmV1Client
	gnsTsmV1          *gnstsmv1.GnsTsmV1Client
	policypkgTsmV1    *policypkgtsmv1.PolicypkgTsmV1Client
	rootTsmV1         *roottsmv1.RootTsmV1Client
	servicegroupTsmV1 *servicegrouptsmv1.ServicegroupTsmV1Client
}

// ConfigTsmV1 retrieves the ConfigTsmV1Client
func (c *Clientset) ConfigTsmV1() configtsmv1.ConfigTsmV1Interface {
	return c.configTsmV1
}

// GnsTsmV1 retrieves the GnsTsmV1Client
func (c *Clientset) GnsTsmV1() gnstsmv1.GnsTsmV1Interface {
	return c.gnsTsmV1
}

// PolicypkgTsmV1 retrieves the PolicypkgTsmV1Client
func (c *Clientset) PolicypkgTsmV1() policypkgtsmv1.PolicypkgTsmV1Interface {
	return c.policypkgTsmV1
}

// RootTsmV1 retrieves the RootTsmV1Client
func (c *Clientset) RootTsmV1() roottsmv1.RootTsmV1Interface {
	return c.rootTsmV1
}

// ServicegroupTsmV1 retrieves the ServicegroupTsmV1Client
func (c *Clientset) ServicegroupTsmV1() servicegrouptsmv1.ServicegroupTsmV1Interface {
	return c.servicegroupTsmV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.configTsmV1, err = configtsmv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.gnsTsmV1, err = gnstsmv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.policypkgTsmV1, err = policypkgtsmv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rootTsmV1, err = roottsmv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.servicegroupTsmV1, err = servicegrouptsmv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.configTsmV1 = configtsmv1.NewForConfigOrDie(c)
	cs.gnsTsmV1 = gnstsmv1.NewForConfigOrDie(c)
	cs.policypkgTsmV1 = policypkgtsmv1.NewForConfigOrDie(c)
	cs.rootTsmV1 = roottsmv1.NewForConfigOrDie(c)
	cs.servicegroupTsmV1 = servicegrouptsmv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.configTsmV1 = configtsmv1.New(c)
	cs.gnsTsmV1 = gnstsmv1.New(c)
	cs.policypkgTsmV1 = policypkgtsmv1.New(c)
	cs.rootTsmV1 = roottsmv1.New(c)
	cs.servicegroupTsmV1 = servicegrouptsmv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
