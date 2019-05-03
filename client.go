package vrlcmsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func (c *ApiClient) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c *ApiClient) checkResponseError(r apiResponse) error {
	if r.statusCode >= 200 && r.statusCode < 400 {
		return nil
	}

	// TODO Add parsing of response
	return fmt.Errorf("received error response code %d", r.statusCode)
}

type apiResponse struct {
	body       io.ReadCloser
	header     http.Header
	statusCode int
	reqURL     *url.URL
}
