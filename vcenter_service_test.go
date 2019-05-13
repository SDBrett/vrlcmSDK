package vrlcmsdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sdbrett/vrlcmsdk/types"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// Test vCenter server object creation for vCenter API service
func TestVcenterAPIService_Create(t *testing.T) {

	var ctx = context.Background()
	expectedURL := "127.0.0.1/lcm/api/v1/action/add/vc"

	t.Run("Test successful vCenter creation", func(t *testing.T) {

		serverResponse := `{"id":"2c551b23da979e75588cb94147e9a","type":"VC_DATA_COLLECTION","state":null,"status":"SUBMITTED","isRetriable":null,"retryParameters":null}`

		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		newVcenter := types.Vcenter{Name: "TestDataCenter", Username: "admin@local.com", Password: "pass", DatacenterName: "dc1", Type: 3,}
		c := NewApiClient("127.0.0.1", true, cli)

		request, err := c.VcenterService.Create(ctx, &newVcenter)
		if err != nil {
			t.Errorf("expected no error when creating datacenter")
		}
		if request.ID != "2c551b23da979e75588cb94147e9a" {
			t.Errorf("expected ID 2c551b23da979e75588cb94147e9a, instead got %s", request.ID)
		}
	})


	t.Run("Test bad JSON decoding", func(t *testing.T) {

		serverResponse := `{"message":"A vCenter with the provided name is already attached to this data center. Use PATCH operation to update the vCenter details.","messageId":"12706","statusCode":400,"documentKind":"com:vmware:xenon:common:ServiceErrorResponse","errorCode":0}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		newVcenter := types.Vcenter{Name: "TestDataCenter", Username: "admin@local.com", Password: "pass", DatacenterName: "dc1", Type: 3,}
		c := NewApiClient("127.0.0.1", true, cli)

		_, err := c.VcenterService.Create(ctx, &newVcenter)
		if err == nil {
			t.Errorf("expected error to be thrown when with server error response")
		}
	})

	t.Run("Test server response error for vCenter creation", func(t *testing.T) {

		serverResponse := "TEST"
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		newVcenter := types.Vcenter{Name: "TestDataCenter", Username: "admin@local.com", Password: "pass", DatacenterName: "dc1", Type: 3,}
		c := NewApiClient("127.0.0.1", true, cli)

		_, err := c.VcenterService.Create(ctx, &newVcenter)
		if err == nil {
			t.Errorf("expected error to be thrown when decoding fails")
		}
	})

}
