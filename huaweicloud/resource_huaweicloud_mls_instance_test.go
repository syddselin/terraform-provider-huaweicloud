// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://www.github.com/huaweicloud/magic-modules
//
// ----------------------------------------------------------------------------

package huaweicloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/huaweicloud/golangsdk"
)

func TestAccMlsInstance_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckMrs(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMlsInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMlsInstance_basic(acctest.RandString(10)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMlsInstanceExists(),
				),
			},
		},
	})
}

func testAccMlsInstance_basic(val string) string {
	return fmt.Sprintf(`
resource "huaweicloud_mrs_cluster_v1" "cluster1" {
  cluster_name = "mrs-cluster-acc%s"
  region = "en-OS_REGION_NAME"
  billing_type = 12
  master_node_num = 2
  core_node_num = 3
  master_node_size = "s1.4xlarge.linux.mrs"
  core_node_size = "s1.xlarge.linux.mrs"
  available_zone_id = "%s"
  vpc_id = "%s"
  subnet_id = "%s"
  cluster_version = "MRS 1.3.0"
  volume_type = "SATA"
  volume_size = 100
  safe_mode = 0
  cluster_type = 0
  node_public_cert_name = "KeyPair-ci"
  cluster_admin_secret = ""
  component_list {
    component_name = "Hadoop"
  }
  component_list {
    component_name = "Spark"
  }
  component_list {
    component_name = "Hive"
  }
  timeouts {
    create = "60m"
  }
}

resource "huaweicloud_mls_instance" "instance" {
  name = "terraform-mls-instancei%s"
  version = "1.5.0"
  flavor = "mls.c2.2xlarge.common"
  network {
    vpc_id = "%s"
    network_id = "%s"
    available_zone = "%s"
    public_ip {
      bind_type = "not_use"
    }
  }
  mrs_cluster {
    id = "${huaweicloud_mrs_cluster_v1.cluster1.id}"
  }

  timeouts {
    create = "60m"
  }
}
	`, val, OS_AVAILABILITY_ZONE, OS_VPC_ID, OS_NETWORK_ID, val, OS_VPC_ID, OS_NETWORK_ID, OS_AVAILABILITY_ZONE)
}

func testAccCheckMlsInstanceDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.sdkClient(OS_REGION_NAME, "mls", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "huaweicloud_mls_instance" {
			continue
		}

		url, err := replaceVarsForTest(rs, "instances/{id}")
		if err != nil {
			return err
		}
		url = client.ServiceURL(url)

		_, err = client.Get(
			url, nil,
			&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}})
		if err == nil {
			return fmt.Errorf("huaweicloud_mls_instance still exists at %s", url)
		}
	}

	return nil
}

func testAccCheckMlsInstanceExists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)
		client, err := config.sdkClient(OS_REGION_NAME, "mls", serviceProjectLevel)
		if err != nil {
			return fmt.Errorf("Error creating sdk client, err=%s", err)
		}

		rs, ok := s.RootModule().Resources["huaweicloud_mls_instance.instance"]
		if !ok {
			return fmt.Errorf("Error checking huaweicloud_mls_instance.instance exist, err=not found huaweicloud_mls_instance.instance")
		}

		url, err := replaceVarsForTest(rs, "instances/{id}")
		if err != nil {
			return fmt.Errorf("Error checking huaweicloud_mls_instance.instance exist, err=building url failed: %s", err)
		}
		url = client.ServiceURL(url)

		_, err = client.Get(
			url, nil,
			&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}})
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return fmt.Errorf("huaweicloud_mls_instance.instance is not exist")
			}
			return fmt.Errorf("Error checking huaweicloud_mls_instance.instance exist, err=send request failed: %s", err)
		}
		return nil
	}
}
