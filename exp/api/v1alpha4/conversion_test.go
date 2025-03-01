/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha4

import (
	"testing"

	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"

	"sigs.k8s.io/cluster-api-provider-azure/exp/api/v1beta1"
)

func TestFuzzyConversion(t *testing.T) {
	g := NewWithT(t)
	scheme := runtime.NewScheme()
	g.Expect(AddToScheme(scheme)).To(Succeed())
	g.Expect(v1beta1.AddToScheme(scheme)).To(Succeed())

	t.Run("for AzureMachinePool", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Scheme: scheme,
		Hub:    &v1beta1.AzureMachinePool{},
		Spoke:  &AzureMachinePool{},
	}))

	t.Run("for AzureManagedCluster", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Scheme: scheme,
		Hub:    &v1beta1.AzureManagedCluster{},
		Spoke:  &AzureManagedCluster{},
	}))

	t.Run("for AzureManagedControlPlane", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Scheme: scheme,
		Hub:    &v1beta1.AzureManagedControlPlane{},
		Spoke:  &AzureManagedControlPlane{},
	}))

	t.Run("for AzureManagedMachinePool", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Scheme: scheme,
		Hub:    &v1beta1.AzureManagedMachinePool{},
		Spoke:  &AzureManagedMachinePool{},
	}))

	t.Run("for AzureMachinePoolMachine", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Scheme: scheme,
		Hub:    &v1beta1.AzureMachinePoolMachine{},
		Spoke:  &AzureMachinePoolMachine{},
	}))

}
