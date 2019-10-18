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
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/huaweicloud/golangsdk"
)

func resourceCsPeeringConnectV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceCsPeeringConnectV1Create,
		Read:   resourceCsPeeringConnectV1Read,
		Delete: resourceCsPeeringConnectV1Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"target_vpc_info": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"project_id": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceCsPeeringConnectV1UserInputParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"terraform_resource_data": d,
		"cluster_id":              d.Get("cluster_id"),
		"name":                    d.Get("name"),
		"target_vpc_info":         d.Get("target_vpc_info"),
	}
}

func resourceCsPeeringConnectV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "cs", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	opts := resourceCsPeeringConnectV1UserInputParams(d)

	arrayIndex := map[string]int{
		"target_vpc_info": 0,
	}

	params, err := buildCsPeeringConnectV1CreateParameters(opts, arrayIndex)
	if err != nil {
		return fmt.Errorf("Error building the request body of api(create), err=%s", err)
	}
	r, err := sendCsPeeringConnectV1CreateRequest(d, params, client)
	if err != nil {
		return fmt.Errorf("Error creating CsPeeringConnectV1, err=%s", err)
	}

	client, err = config.sdkClient(GetRegion(d, config), "network", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	err = actionCsPeeringConnectV1AcceptPeering(d, r, client)
	if err != nil {
		return err
	}

	timeout := d.Timeout(schema.TimeoutCreate)

	client, err = config.sdkClient(GetRegion(d, config), "cs", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating async sdk client, err=%s", err)
	}

	obj, err := asyncWaitCsPeeringConnectV1Create(d, config, r, client, timeout)
	if err != nil {
		return err
	}
	id, err := navigateValue(obj, []string{"peering", "id"}, nil)
	if err != nil {
		return fmt.Errorf("Error constructing id, err=%s", err)
	}
	d.SetId(convertToStr(id))

	return resourceCsPeeringConnectV1Read(d, meta)
}

func resourceCsPeeringConnectV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "cs", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	res := make(map[string]interface{})

	v, err := sendCsPeeringConnectV1ReadRequest(d, client)
	if err != nil {
		return err
	}
	res["read"] = fillCsPeeringConnectV1ReadRespBody(v)

	states, err := flattenCsPeeringConnectV1Options(res)
	if err != nil {
		return err
	}

	return setCsPeeringConnectV1States(d, states)
}

func resourceCsPeeringConnectV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "cs", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	url, err := replaceVars(d, "reserved_cluster/{cluster_id}/peering/{id}", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	log.Printf("[DEBUG] Deleting PeeringConnect %q", d.Id())
	r := golangsdk.Result{}
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		OkCodes:      successHTTPCodes,
		JSONBody:     nil,
		JSONResponse: nil,
		MoreHeaders:  map[string]string{"Content-Type": "application/json"},
	})
	if r.Err != nil {
		return fmt.Errorf("Error deleting PeeringConnect %q, err=%s", d.Id(), r.Err)
	}

	_, err = asyncWaitCsPeeringConnectV1Delete(d, config, r.Body, client, d.Timeout(schema.TimeoutDelete))
	return err
}

func buildCsPeeringConnectV1CreateParameters(opts map[string]interface{}, arrayIndex map[string]int) (interface{}, error) {
	params := make(map[string]interface{})

	v, err := navigateValue(opts, []string{"target_vpc_info", "project_id"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["accept_tenant_id"] = v
	}

	v, err = navigateValue(opts, []string{"name"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["name"] = v
	}

	v, err = navigateValue(opts, []string{"target_vpc_info", "vpc_id"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["vpc_id"] = v
	}

	return params, nil
}

func sendCsPeeringConnectV1CreateRequest(d *schema.ResourceData, params interface{},
	client *golangsdk.ServiceClient) (interface{}, error) {
	url, err := replaceVars(d, "reserved_cluster/{cluster_id}/peering", nil)
	if err != nil {
		return nil, err
	}
	url = client.ServiceURL(url)

	r := golangsdk.Result{}
	_, r.Err = client.Post(url, params, &r.Body, &golangsdk.RequestOpts{
		OkCodes: successHTTPCodes,
	})
	if r.Err != nil {
		return nil, fmt.Errorf("Error running api(create), err=%s", r.Err)
	}
	return r.Body, nil
}

func asyncWaitCsPeeringConnectV1Create(d *schema.ResourceData, config *Config, result interface{},
	client *golangsdk.ServiceClient, timeout time.Duration) (interface{}, error) {

	data := make(map[string]interface{})
	pathParameters := map[string][]string{
		"peering_id": []string{"peering", "id"},
	}
	for key, path := range pathParameters {
		value, err := navigateValue(result, path, nil)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving async operation path parameter, err=%s", err)
		}
		data[key] = value
	}

	url, err := replaceVars(d, "reserved_cluster/{cluster_id}/peering/{peering_id}", data)
	if err != nil {
		return nil, err
	}
	url = client.ServiceURL(url)

	return waitToFinish(
		[]string{"ACTIVE"},
		[]string{"PENDING_ACCEPTANCE"},
		timeout, 1*time.Second,
		func() (interface{}, string, error) {
			r := golangsdk.Result{}
			_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
				MoreHeaders: map[string]string{"Content-Type": "application/json"}})
			if r.Err != nil {
				return nil, "", nil
			}

			status, err := navigateValue(r.Body, []string{"peering", "status"}, nil)
			if err != nil {
				return nil, "", nil
			}
			return r.Body, convertToStr(status), nil
		},
	)
}

func asyncWaitCsPeeringConnectV1Delete(d *schema.ResourceData, config *Config, result interface{},
	client *golangsdk.ServiceClient, timeout time.Duration) (interface{}, error) {

	url, err := replaceVars(d, "reserved_cluster/{cluster_id}/peering/{id}", nil)
	if err != nil {
		return nil, err
	}
	url = client.ServiceURL(url)

	return waitToFinish(
		[]string{"Done"}, []string{"Pending"}, timeout, 1*time.Second,
		func() (interface{}, string, error) {
			r := golangsdk.Result{}
			_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
				OkCodes:     []int{200, 400},
				MoreHeaders: map[string]string{"Content-Type": "application/json"}})
			if r.Err != nil {
				return nil, "", nil
			}

			if checkCsPeeringConnectV1DeleteFinished(r.Body) {
				return r.Body, "Done", nil
			}
			return r.Body, "Pending", nil
		},
	)
}

func sendCsPeeringConnectV1ReadRequest(d *schema.ResourceData, client *golangsdk.ServiceClient) (interface{}, error) {
	url, err := replaceVars(d, "reserved_cluster/{cluster_id}/peering/{id}", nil)
	if err != nil {
		return nil, err
	}
	url = client.ServiceURL(url)

	r := golangsdk.Result{}
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"}})
	if r.Err != nil {
		return nil, fmt.Errorf("Error running api(read) for resource(CsPeeringConnectV1), err=%s", r.Err)
	}

	v, err := navigateValue(r.Body, []string{"peering"}, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func fillCsPeeringConnectV1ReadRespBody(body interface{}) interface{} {
	result := make(map[string]interface{})
	val, ok := body.(map[string]interface{})
	if !ok {
		val = make(map[string]interface{})
	}

	if v, ok := val["accept_vpc_info"]; ok {
		result["accept_vpc_info"] = fillCsPeeringConnectV1ReadRespAcceptVpcInfo(v)
	} else {
		result["accept_vpc_info"] = nil
	}

	if v, ok := val["name"]; ok {
		result["name"] = v
	} else {
		result["name"] = nil
	}

	return result
}

func fillCsPeeringConnectV1ReadRespAcceptVpcInfo(value interface{}) interface{} {
	if value == nil {
		return nil
	}

	value1, ok := value.(map[string]interface{})
	if !ok {
		value1 = make(map[string]interface{})
	}
	result := make(map[string]interface{})

	if v, ok := value1["tenant_id"]; ok {
		result["tenant_id"] = v
	} else {
		result["tenant_id"] = nil
	}

	if v, ok := value1["vpc_id"]; ok {
		result["vpc_id"] = v
	} else {
		result["vpc_id"] = nil
	}

	return result
}

func flattenCsPeeringConnectV1Options(response map[string]interface{}) (map[string]interface{}, error) {
	opts := make(map[string]interface{})

	v, err := navigateValue(response, []string{"read", "name"}, nil)
	if err != nil {
		return nil, fmt.Errorf("Error flattening PeeringConnect:name, err: %s", err)
	}
	opts["name"] = v

	v, err = flattenCsPeeringConnectV1TargetVpcInfo(response, nil)
	if err != nil {
		return nil, fmt.Errorf("Error flattening PeeringConnect:target_vpc_info, err: %s", err)
	}
	opts["target_vpc_info"] = v

	return opts, nil
}

func flattenCsPeeringConnectV1TargetVpcInfo(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	r := make(map[string]interface{})

	v, err := navigateValue(d, []string{"read", "accept_vpc_info", "tenant_id"}, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error flattening PeeringConnect:project_id, err: %s", err)
	}
	r["project_id"] = v

	v, err = navigateValue(d, []string{"read", "accept_vpc_info", "vpc_id"}, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error flattening PeeringConnect:vpc_id, err: %s", err)
	}
	r["vpc_id"] = v

	for _, v := range r {
		if v != nil {
			return []interface{}{r}, nil
		}
	}
	return nil, nil
}

func setCsPeeringConnectV1States(d *schema.ResourceData, opts map[string]interface{}) error {
	for k, v := range opts {
		if err := d.Set(k, v); err != nil {
			return fmt.Errorf("Error setting CsPeeringConnectV1:%s, err: %s", k, err)
		}
	}
	return nil
}

func actionCsPeeringConnectV1AcceptPeering(d *schema.ResourceData, result interface{}, client *golangsdk.ServiceClient) error {
	pathParameters := map[string][]string{
		"peering_id": []string{"peering", "id"},
	}
	var data = make(map[string]interface{})
	for key, path := range pathParameters {
		value, err := navigateValue(result, path, nil)
		if err != nil {
			return fmt.Errorf("Error retrieving path parameter, err=%s", err)
		}
		data[key] = value
	}
	url, err := replaceVars(d, "v2.0/vpc/peerings/{peering_id}/accept", data)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	r := golangsdk.Result{}
	_, r.Err = client.Put(url, nil, nil, &golangsdk.RequestOpts{
		OkCodes: successHTTPCodes,
	})
	if r.Err != nil {
		return fmt.Errorf("Error run action of accept_peering, err=%s", r.Err)
	}
	return nil
}
