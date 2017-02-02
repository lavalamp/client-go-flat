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

package meta

import (
	"github.com/lavalamp/client-go-flat/apimachinery/pkg/apis/meta/v1/unstructured"
	"github.com/lavalamp/client-go-flat/apimachinery/pkg/runtime/schema"
)

// InterfacesForUnstructured returns VersionInterfaces suitable for
// dealing with unstructured.Unstructured objects.
func InterfacesForUnstructured(schema.GroupVersion) (*VersionInterfaces, error) {
	return &VersionInterfaces{
		ObjectConvertor:  &unstructured.UnstructuredObjectConverter{},
		MetadataAccessor: NewAccessor(),
	}, nil
}