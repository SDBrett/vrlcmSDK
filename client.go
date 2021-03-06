package vrlcmsdk

import (
	"net/http"
)

// API client for interaction with LCM
type ApiClient struct {
	basePath      string
	cmsPath       string
	Token         string
	AllowInsecure bool
	Host          string
	httpClient    *http.Client
	headers       *http.Header

	DatacenterService *DatacenterAPIService
	VcenterService    *VcenterAPIService
	SettingsService *SettingsAPIService
}

// Defines services
type service struct {
	client *ApiClient
}

// Create new API client
func NewApiClient(host string, ignoreSSL bool, httpClient *http.Client) ApiClient {

	c := &ApiClient{
		AllowInsecure: ignoreSSL,
		Host:          host,
		basePath:      host + "/lcm/api",
		cmsPath:       host + "/cms/api",
	}

	if httpClient == nil {
		c.newDefaultClient()
	} else {
		c.httpClient = httpClient
	}

	c.SetDefaultHeaders()
	c.DatacenterService = &DatacenterAPIService{client: c}
	c.VcenterService = &VcenterAPIService{client: c}
	c.SettingsService = &SettingsAPIService{client: c}
	return *c
}
