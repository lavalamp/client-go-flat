/*
Copyright 2016 The Kubernetes Authors.

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

package kubernetes

// These imports are the API groups the client will support.
import (
	"fmt"

	"github.com/lavalamp/client-go-flat/pkg/api"
	_ "github.com/lavalamp/client-go-flat/pkg/api/install"
	_ "github.com/lavalamp/client-go-flat/pkg/apis/apps/install"
	_ "github.com/lavalamp/client-go-flat/pkg/apis/authentication/install"
	_ "github.com/lavalamp/client-go-flat/pkg/apis/authorization/install"
	_ "github.com/lavalamp/client-go-flat/pkg/apis/autoscaling/install"
	_ "github.com/lavalamp/client-go-flat/pkg/apis/batch/install"
	_ "github.com/lavalamp/client-go-flat/pkg/apis/certificates/install"
	_ "github.com/lavalamp/client-go-flat/pkg/apis/extensions/install"
	_ "github.com/lavalamp/client-go-flat/pkg/apis/policy/install"
	_ "github.com/lavalamp/client-go-flat/pkg/apis/rbac/install"
	_ "github.com/lavalamp/client-go-flat/pkg/apis/storage/install"
)

func init() {
	if missingVersions := api.Registry.ValidateEnvRequestedVersions(); len(missingVersions) != 0 {
		panic(fmt.Sprintf("KUBE_API_VERSIONS contains versions that are not installed: %q.", missingVersions))
	}
}
