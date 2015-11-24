package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func DiscoveryResource() *schema.Resource {
	return &schema.Resource{
		Create: createToken,
		Read:   readToken,
		Update: createToken,
		Delete: deleteToken,
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://discovery.etcd.io/new",
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func createToken(d *schema.ResourceData, m interface{}) error {
	u, err := url.Parse(d.Get("address").(string))
	if err != nil {
		return err
	}
	if size, ok := d.GetOk("size"); ok {
		q := make(url.Values)
		q.Set("size", strconv.Itoa(size.(int)))
		u.RawQuery = q.Encode()
	}
	r, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	d.SetId(string(body))
	return nil
}

func readToken(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteToken(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
