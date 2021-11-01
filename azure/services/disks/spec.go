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

// DiskSpec defines the specification for a Disk.
type DiskSpec struct {
	Name string
}

// ResourceName returns the name of the disk.
func (s *DiskSpec) ResourceName() string {
	return s.Name
}

// OwnerResourceName is a no-op for disks.
func (s *DiskSpec) OwnerResourceName() string {
	// TODO(karuppiah7890): Check if disks have any owner resource name
	return ""
}

// ResourceGroupName returns the name of the resource group the disk is in.
func (s *DiskSpec) ResourceGroupName() string {
	// TODO(karuppiah7890): Implement return of resource group name from spec field
	return ""
}

// Parameters returns the parameters for the route table.
func (s *DiskSpec) Parameters(existing interface{}) (interface{}, error) {
	// TODO(karuppiah7890): Implement returning of params
	return nil, nil
}
