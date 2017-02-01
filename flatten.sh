#!/bin/bash

# Copyright YEAR The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


mv vendor/k8s.io/apimachinery/ apimachinery

find . | grep "[.]go$" | xargs -n 1 sed -i 's|k8s.io/apimachinery|k8s.io/client-go/apimachinery|'
find . | grep "[.]go$" | xargs -n 1 sed -i 's|k8s.io/client-go|github.com/lavalamp/client-go-flat|'

git add apimachinery
git commit -a -m "automated dependency flattening"
