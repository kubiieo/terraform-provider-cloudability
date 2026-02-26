package cloudability

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kubiieo/cloudability-sdk-go/cloudability"
)

// Config - Provider Config
type Config struct {
	APIKey string
	Region string
}

// NewConfig - Return a new Config instance
func NewConfig(d *schema.ResourceData) *Config {
	c := &Config{
		APIKey: d.Get("apikey").(string),
		Region: d.Get("region").(string),
	}
	return c
}

// Client - Retrun a new Cloudabiity Client instance
func (c *Config) Client() *cloudability.Client {
	client := cloudability.NewClient(c.APIKey)
	return client
}
