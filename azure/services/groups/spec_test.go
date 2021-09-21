package groups_test

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha4"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/groups"
)

func TestParameters(t *testing.T) {
	t.Run("when resource group already exists and additional tags has been updated in group spec return resource group with updated tags", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name                      string
			existingResourceGroupTags map[string]*string
			groupSpecTags             infrav1.Tags
			expectedResourceGroupTags map[string]*string
		}{
			{
				name: "new tag with value has been added",
				existingResourceGroupTags: map[string]*string{
					"Name": to.StringPtr("test-group"),
					"sigs.k8s.io_cluster-api-provider-azure_role":                 to.StringPtr("common"),
					"sigs.k8s.io_cluster-api-provider-azure_cluster_test-cluster": to.StringPtr("owned"),
				},
				groupSpecTags: infrav1.Tags{"environment": "dev"},
				expectedResourceGroupTags: map[string]*string{
					"Name": to.StringPtr("test-group"),
					"sigs.k8s.io_cluster-api-provider-azure_role":                 to.StringPtr("common"),
					"sigs.k8s.io_cluster-api-provider-azure_cluster_test-cluster": to.StringPtr("owned"),
					"environment": to.StringPtr("dev"),
				},
			},
			{
				name: "existing tag with updated value",
				existingResourceGroupTags: map[string]*string{
					"Name": to.StringPtr("test-group"),
					"sigs.k8s.io_cluster-api-provider-azure_role":                 to.StringPtr("common"),
					"sigs.k8s.io_cluster-api-provider-azure_cluster_test-cluster": to.StringPtr("owned"),
					"environment": to.StringPtr("staging"),
				},
				groupSpecTags: infrav1.Tags{"environment": "dev"},
				expectedResourceGroupTags: map[string]*string{
					"Name": to.StringPtr("test-group"),
					"sigs.k8s.io_cluster-api-provider-azure_role":                 to.StringPtr("common"),
					"sigs.k8s.io_cluster-api-provider-azure_cluster_test-cluster": to.StringPtr("owned"),
					"environment": to.StringPtr("dev"),
				},
			},
			{
				name: "existing tag with nil value has been updated",
				existingResourceGroupTags: map[string]*string{
					"Name": to.StringPtr("test-group"),
					"sigs.k8s.io_cluster-api-provider-azure_role":                 to.StringPtr("common"),
					"sigs.k8s.io_cluster-api-provider-azure_cluster_test-cluster": to.StringPtr("owned"),
					"environment": nil,
				},
				groupSpecTags: infrav1.Tags{"environment": "dev"},
				expectedResourceGroupTags: map[string]*string{
					"Name": to.StringPtr("test-group"),
					"sigs.k8s.io_cluster-api-provider-azure_role":                 to.StringPtr("common"),
					"sigs.k8s.io_cluster-api-provider-azure_cluster_test-cluster": to.StringPtr("owned"),
					"environment": to.StringPtr("dev"),
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()
				g := NewWithT(t)
				resourceGroupName := "test-group"
				resourceGroupLocation := "test-location"

				existingResourceGroup := resources.Group{
					Name:     to.StringPtr(resourceGroupName),
					Location: to.StringPtr(resourceGroupLocation),
					Tags:     testCase.existingResourceGroupTags,
				}

				groupSpec := groups.GroupSpec{
					Name:           resourceGroupName,
					Location:       resourceGroupLocation,
					ClusterName:    "test-cluster",
					AdditionalTags: testCase.groupSpecTags,
				}

				expectedResourceGroup := resources.Group{
					Location: to.StringPtr(resourceGroupLocation),
					Tags:     testCase.expectedResourceGroupTags,
				}

				resourceGroup, err := groupSpec.Parameters(existingResourceGroup)

				g.Expect(err).NotTo(HaveOccurred())

				g.Expect(resourceGroup).To(Equal(expectedResourceGroup))
			})
		}
	})

	t.Run("when resource group already exists and additional tags have not been updated in group spec return nil", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name                      string
			existingResourceGroupTags map[string]*string
			groupSpecTags             infrav1.Tags
		}{
			{
				name: "same set of additional tags with tag values",
				existingResourceGroupTags: map[string]*string{
					"Name": to.StringPtr("test-group"),
					"sigs.k8s.io_cluster-api-provider-azure_role":                 to.StringPtr("common"),
					"sigs.k8s.io_cluster-api-provider-azure_cluster_test-cluster": to.StringPtr("owned"),
					"environment": to.StringPtr("dev"),
				},
				groupSpecTags: infrav1.Tags{"environment": "dev"},
			},
			{
				name: "no additional tags present in spec",
				existingResourceGroupTags: map[string]*string{
					"Name": to.StringPtr("test-group"),
					"sigs.k8s.io_cluster-api-provider-azure_role":                 to.StringPtr("common"),
					"sigs.k8s.io_cluster-api-provider-azure_cluster_test-cluster": to.StringPtr("owned"),
				},
				groupSpecTags: infrav1.Tags{},
			},
			{
				name: "nil additional tags in spec",
				existingResourceGroupTags: map[string]*string{
					"Name": to.StringPtr("test-group"),
					"sigs.k8s.io_cluster-api-provider-azure_role":                 to.StringPtr("common"),
					"sigs.k8s.io_cluster-api-provider-azure_cluster_test-cluster": to.StringPtr("owned"),
				},
				groupSpecTags: nil,
			},
			{
				name: "no additional tags present in spec with unmanaged tags in existing resource group",
				existingResourceGroupTags: map[string]*string{
					"Name": to.StringPtr("test-group"),
					"sigs.k8s.io_cluster-api-provider-azure_role":                 to.StringPtr("common"),
					"sigs.k8s.io_cluster-api-provider-azure_cluster_test-cluster": to.StringPtr("owned"),
					// TODO: Check if this extra unmanaged tag can exist
					"environment": to.StringPtr("dev"),
				},
				groupSpecTags: infrav1.Tags{},
			},
			{
				name: "nil additional tags in spec with unmanaged tags in existing resource group",
				existingResourceGroupTags: map[string]*string{
					"Name": to.StringPtr("test-group"),
					"sigs.k8s.io_cluster-api-provider-azure_role":                 to.StringPtr("common"),
					"sigs.k8s.io_cluster-api-provider-azure_cluster_test-cluster": to.StringPtr("owned"),
					// TODO: Check if this extra unmanaged tag can exist
					"environment": to.StringPtr("dev"),
				},
				groupSpecTags: nil,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				g := NewWithT(t)
				resourceGroupName := "test-group"
				resourceGroupLocation := "test-location"

				existingResourceGroup := resources.Group{
					Name:     to.StringPtr(resourceGroupName),
					Location: to.StringPtr(resourceGroupLocation),
					Tags:     testCase.existingResourceGroupTags,
				}

				groupSpec := groups.GroupSpec{
					Name:           resourceGroupName,
					Location:       resourceGroupLocation,
					ClusterName:    "test-cluster",
					AdditionalTags: testCase.groupSpecTags,
				}

				resourceGroup, err := groupSpec.Parameters(existingResourceGroup)

				g.Expect(err).NotTo(HaveOccurred())

				g.Expect(resourceGroup).To(BeNil())
			})
		}
	})
}
