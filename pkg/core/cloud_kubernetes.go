/*
Copyright 2021 The Pixiu Authors.

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

package core

import (
	pixiukubernetes "github.com/caoyingjunz/gopixiu/pkg/core/kubernetes"
)

func (c *cloud) Namespaces(cloud string) pixiukubernetes.NamespaceInterface {
	return pixiukubernetes.NewNamespaces(clientSets.Get(cloud), cloud)
}

func (c *cloud) Services(cloud string) pixiukubernetes.ServiceInterface {
	return pixiukubernetes.NewServices(clientSets.Get(cloud), cloud)
}

func (c *cloud) Deployments(cloud string) pixiukubernetes.DeploymentInterface {
	return pixiukubernetes.NewDeployments(clientSets.Get(cloud), cloud)
}

func (c *cloud) StatefulSets(cloud string) pixiukubernetes.StatefulSetInterface {
	return pixiukubernetes.NewStatefulSets(clientSets.Get(cloud), cloud)
}

func (c *cloud) Jobs(cloud string) pixiukubernetes.JobInterface {
	return pixiukubernetes.NewJobs(clientSets.Get(cloud), cloud)
}