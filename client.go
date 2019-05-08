package vrlcmsdk

import (
	"net/http"
)

type ApiClient struct {
	basePath      string
	cmsPath       string
	token         string
	AllowInsecure bool
	Host          string
	httpClient    *http.Client
	headers       *http.Header

	DatacenterService *DatacenterAPIService
}

type service struct {
	client *ApiClient
}

func NewApiClient(host string, ignoreSSL bool, httpClient *http.Client) ApiClient {

	//var c ApiClient
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
