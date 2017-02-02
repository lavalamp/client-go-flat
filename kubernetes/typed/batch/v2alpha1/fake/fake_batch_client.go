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

package fake

import (
	v2alpha1 "github.com/lavalamp/client-go-flat/kubernetes/typed/batch/v2alpha1"
	rest "github.com/lavalamp/client-go-flat/rest"
	testing "github.com/lavalamp/client-go-flat/testing"
)

type FakeBatchV2alpha1 struct {
	*testing.Fake
}

func (c *FakeBatchV2alpha1) CronJobs(namespace string) v2alpha1.CronJobInterface {
	return &FakeCronJobs{c, namespace}
}

func (c *FakeBatchV2alpha1) Jobs(namespace string) v2alpha1.JobInterface {
	return &FakeJobs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBatchV2alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
