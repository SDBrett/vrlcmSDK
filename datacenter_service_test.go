package vrlcmsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sdbrett/vrlcmsdk/types"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func getMockDatacenteres() []types.Datacenter {

	var dcArray []types.Datacenter

	dcbyte := []byte(`[{"id":"2c551b23da979e755885a0b8be910","datacenterName":"TEST"},{"id":"2c551b23da979e755885ab04ef11b","datacenterName":"sdkDC"},{"id":"2c551b23da979e755885ac51acb58","datacenterName":"sdkD32"},{"id":"2c551b23da979e75587c68ad7dfda","datacenterName":"test3"},{"id":"2c551b23da979e75584db8aaa9688","datacenterName":"test"},{"id":"2c551b23da979e75587c936176f3a","datacenterName":"mytest"},{"id":"2c551b23da979e7558845e525b971","datacenterName":"APITEST2"},{"id":"2c551b23da979e75587c66266a612","datacenterName":"Test2"},{"id":"2c551b23da979e7558845e2cff610","datacenterName":"APITEST1"},{"id":"2c551b23da979e755885a8e3efc91","datacenterName":"TES232T"},{"id":"2c551b23da979e755885ac2553318","datacenterName":"sdkDC2"},{"id":"2c551b23da979e75584f634e35ff8","datacenterName":"BANGALORE_DATA_CENTER"}]`)

	err := json.Unmarshal(dcbyte, &dcArray)
	if err != nil {
		log.Println("error marshalling datacenters")
	}

	return dcArray
}

func TestDatacenterAPIService_GetDatacenter(t *testing.T) {

	var ctx = context.Background()

	expectedURL := "127.0.0.1/lcm/api/v1/view/datacenter"
	allDatacentersReponse := `[{"id":"2c551b23da979e755885a0b8be910","name":"TEST"},{"id":"2c551b23da979e755885ab04ef11b","name":"sdkDC"},{"id":"2c551b23da979e755885ac51acb58","name":"sdkD32"},{"id":"2c551b23da979e75587c68ad7dfda","name":"test3"},{"id":"2c551b23da979e75584db8aaa9688","name":"test"},{"id":"2c551b23da979e75587c936176f3a","name":"mytest"},{"id":"2c551b23da979e7558845e525b971","name":"APITEST2"},{"id":"2c551b23da979e75587c66266a612","name":"Test2"},{"id":"2c551b23da979e7558845e2cff610","name":"APITEST1"},{"id":"2c551b23da979e755885a8e3efc91","name":"TES232T"},{"id":"2c551b23da979e755885ac2553318","name":"sdkDC2"},{"id":"2c551b23da979e75584f634e35ff8","name":"BANGALORE_DATA_CENTER"}]`
	mockDatacenters := getMockDatacenteres()

	t.Run("Test successful get all datacenter", func(t *testing.T) {
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "GET" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}
			dcQuery := req.URL.Query().Get("datacenterId")

			if dcQuery != "" {
				for _, i := range mockDatacenters {
					if i.ID == dcQuery {
						strB, _ := json.Marshal(i)

						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       ioutil.NopCloser(bytes.NewReader([]byte(strB))),
						}, nil
					}
				}
				return &http.Response{
					StatusCode: http.StatusNotFound,
				}, nil
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(allDatacentersReponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		dataCenters, err := c.DatacenterService.GetAllDatacenters(ctx)
		if err != nil {
			t.Error(err)
		}

		if len(dataCenters.Datacenter) != 12 {
			t.Errorf("expected 12 datacenters, received %d", len(dataCenters.Datacenter))
		}

		for dc, v := range dataCenters.Datacenter {
			if v.ID == "" {
				t.Errorf("expected ID field to have value, found empty value at pos %d", dc)
			}
		}
	})

	t.Run("Test error during get individual datacenters", func(t *testing.T) {
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "GET" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}
			dcQuery := req.URL.Query().Get("datacenterId")

			if dcQuery != "" {
				return &http.Response{
					StatusCode: http.StatusNotFound,
				}, nil
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(allDatacentersReponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		_, err := c.DatacenterService.GetAllDatacenters(ctx)
		if err == nil {
			t.Errorf("expected error return getting datacenters, got none")
		}

	})

	t.Run("Test failed get all datacenter", func(t *testing.T) {

		cli := newMockClient(errorMock(http.StatusInternalServerError, "Server error"))
		c := NewApiClient("127.0.0.1", true, cli)

		_, err := c.DatacenterService.GetAllDatacenters(ctx)
		if err == nil {
			t.Error("expected response error to be returned, received no error")
		}
	})
}
