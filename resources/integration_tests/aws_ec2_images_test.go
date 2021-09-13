package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEc2Images(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Images(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Ec2Images().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(
					squirrel.Like{"name": fmt.Sprintf("aws-ec2-images-image-%s%s%%", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"image_type":          "machine",
						"ena_support":         true,
						"hypervisor":          "xen",
						"platform_details":    "Linux/UNIX",
						"public":              false,
						"root_device_name":    "/dev/xvda",
						"root_device_type":    "ebs",
						"sriov_net_support":   "simple",
						"usage_operation":     "RunInstances",
						"virtualization_type": "hvm",
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_ec2_image_block_device_mappings",
					ForeignKeyName: "image_cq_id",
					Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
						return sq.Where(squirrel.Eq{"device_name": "/dev/xvda"})
					},
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"device_name":               "/dev/xvda",
							"ebs_delete_on_termination": true,
							"ebs_encrypted":             false,
							"ebs_volume_size":           float64(8),
							"ebs_volume_type":           "gp2",
						},
					}},
				},
				{
					Name:           "aws_ec2_image_block_device_mappings",
					ForeignKeyName: "image_cq_id",
					Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
						return sq.Where(squirrel.Eq{"device_name": "/dev/xvdb"})
					},
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"device_name":               "/dev/xvdb",
							"ebs_delete_on_termination": true,
							"ebs_encrypted":             false,
							"ebs_volume_size":           float64(20),
							"ebs_volume_type":           "gp2",
						},
					}},
				},
			},
		}
	})
}