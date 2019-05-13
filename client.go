package vrlcmsdk

import (
	"net/http"
)

// API client for interaction with LCM
type ApiClient struct {
	basePath      string
	cmsPath       string
	token         string
	AllowInsecure bool
	Host          string
	httpClient    *http.Client
	headers       *http.Header

	DatacenterService *DatacenterAPIService
	VcenterService *VcenterAPIService
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
		basePath:      host + "/lcm/api/v1",
		cmsPath:       host + "/cms/api/v1",
	}

	if httpClient == nil {
		c.newDefaultClient()
	} else {
		c.httpClient = httpClient
	}

	c.setDefaultHeaders()
	c.DatacenterService = &DatacenterAPIService{client: c}

	return *c
}
