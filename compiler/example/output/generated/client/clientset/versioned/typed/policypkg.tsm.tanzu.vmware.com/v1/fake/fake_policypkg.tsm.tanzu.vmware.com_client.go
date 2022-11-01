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

package fake

import (
	v1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/crd_generatedclient/clientset/versioned/typed/policypkg.tsm.tanzu.vmware.com/v1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePolicypkgTsmV1 struct {
	*testing.Fake
}

func (c *FakePolicypkgTsmV1) ACPConfigs() v1.ACPConfigInterface {
	return &FakeACPConfigs{c}
}

func (c *FakePolicypkgTsmV1) AccessControlPolicies() v1.AccessControlPolicyInterface {
	return &FakeAccessControlPolicies{c}
}

func (c *FakePolicypkgTsmV1) AdditionalPolicyDatas() v1.AdditionalPolicyDataInterface {
	return &FakeAdditionalPolicyDatas{c}
}

func (c *FakePolicypkgTsmV1) RandomPolicyDatas() v1.RandomPolicyDataInterface {
	return &FakeRandomPolicyDatas{c}
}

func (c *FakePolicypkgTsmV1) VMpolicies() v1.VMpolicyInterface {
	return &FakeVMpolicies{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePolicypkgTsmV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
