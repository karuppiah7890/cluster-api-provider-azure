/*
Copyright 2019 The Kubernetes Authors.

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

package disks

import (
	"context"

	"github.com/go-logr/logr"
	// "github.com/pkg/errors"

	"sigs.k8s.io/cluster-api-provider-azure/azure"
	"sigs.k8s.io/cluster-api-provider-azure/util/tele"
)

const serviceName = "disks"

// DiskScope defines the scope interface for a disk service.
type DiskScope interface {
	logr.Logger
	azure.ClusterDescriber
	azure.AsyncStatusUpdater
	DiskSpecs() []azure.ResourceSpecGetter
}

// Service provides operations on Azure resources.
type Service struct {
	Scope DiskScope
	client
}

// New creates a new disks service.
func New(scope DiskScope) *Service {
	return &Service{
		Scope:  scope,
		client: newClient(scope),
	}
}

// Reconcile on disk is currently no-op. OS disks should only be deleted and will create with the VM automatically.
func (s *Service) Reconcile(ctx context.Context) error {
	_, _, done := tele.StartSpanWithLogger(ctx, "disks.Service.Reconcile")
	defer done()

	return nil
}

// Delete deletes the disk associated with a VM.
func (s *Service) Delete(ctx context.Context) error {
	ctx, _, done := tele.StartSpanWithLogger(ctx, "disks.Service.Delete")
	defer done()

	// TODO(karuppiah7890): I'm assuming that all the disk specs from DiskSpecs() are managed disks as they
	// belong to managed VMs. So I'm assuming we don't have to do any checks to see if a disk is managed or not.
	// Verify these assumptions!
	// TODO(karuppiah7890): Implement this part - the for loop with error handling etc along with tests
	// for _, diskSpec := range s.Scope.DiskSpecs() {
	// async.DeleteResource(ctx, s.Scope, s.client, diskSpec, serviceName)
	// }
	return nil
}
