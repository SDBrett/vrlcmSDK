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

func TestDatacenterAPIService_GetAllDatacenters(t *testing.T) {

	var ctx = context.Background()

	expectedURL := "127.0.0.1/lcm/api/v1/view/datacenter"
	allDatacentersReponse := `[{"id":"2c551b23da979e755885a0b8be910","name":"TEST"},{"id":"2c551b23da979e755885ab04ef11b","name":"sdkDC"},{"id":"2c551b23da979e755885ac51acb58","name":"sdkD32"},{"id":"2c551b23da979e75587c68ad7dfda","name":"test3"},{"id":"2c551b23da979e75584db8aaa9688","name":"test"},{"id":"2c551b23da979e75587c936176f3a","name":"mytest"},{"id":"2c551b23da979e7558845e525b971","name":"APITEST2"},{"id":"2c551b23da979e75587c66266a612","name":"Test2"},{"id":"2c551b23da979e7558845e2cff610","name":"APITEST1"},{"id":"2c551b23da979e755885a8e3efc91","name":"TES232T"},{"id":"2c551b23da979e755885ac2553318","name":"sdkDC2"},{"id":"2c551b23da979e75584f634e35ff8","name":"BANGALORE_DATA_CENTER"}]`
	mockDatacenters := getMockDatacenteres()

	t.Run("Test successful get all datacenter", func(t *testing.T) {
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
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

	t.Run("Test decoding error for datacenters", func(t *testing.T) {

		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "GET" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}
			dcQuery := req.URL.Query().Get("datacenterId")

			if dcQuery != "" {
				return &http.Response{
					StatusCode: http.StatusBadGateway,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
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
}

func TestDatacenterAPIService_GetDatacenter(t *testing.T) {

	ctx := context.Background()
	expectedURL := "127.0.0.1/lcm/api/v1/view/datacenter"
	mockDatacenters := getMockDatacenteres()

	t.Run("Test successful get datacenter", func(t *testing.T) {
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
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
				StatusCode: http.StatusNotFound,
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		dataCenter, err := c.DatacenterService.GetDatacenter(ctx, "2c551b23da979e755885a0b8be910")
		if err != nil {
			t.Error(err)
		}

		if dataCenter.ID != "2c551b23da979e755885a0b8be910" {
			t.Errorf("expected ID field to equal 2c551b23da979e755885a0b8be910, instead got %s", dataCenter.ID)
		}
	})

	t.Run("Test decoding error when getting datacenter", func(t *testing.T) {
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "GET" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		_, err := c.DatacenterService.GetDatacenter(ctx, "2c551b23da979e755885a0b8be910")
		if err == nil {
			t.Error("expected to receive decoding error, received no error")
		}
	})
}

func TestDatacenterAPIService_Create(t *testing.T) {

	var ctx = context.Background()
	expectedURL := "127.0.0.1/lcm/api/v1/action/create/datacenter"

	t.Run("Test successful datacenter creation", func(t *testing.T) {

		creationResponse := `{"id":"2c551b23da979e75588bd59d7418b","type":null,"state":null,"status":"SUCCESS","isRetriable":null,"retryParameters":null}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(creationResponse))),
			}, nil
		})

		newDatacenter := types.Datacenter{Name: "TestDataCenter", City: "Melbourne,VIC"}
		c := NewApiClient("127.0.0.1", true, cli)

		err := c.DatacenterService.Create(ctx, &newDatacenter)
		if err != nil {
			t.Errorf("expected no error when creating datacenter")
		}
		if newDatacenter.ID != "2c551b23da979e75588bd59d7418b" {
			t.Errorf("incorrect ID returned")
		}
		if newDatacenter.Name != "TestDataCenter" {
			t.Errorf("expected new datacenter name TestDataCenter, received %s", newDatacenter.Name)
		}
		if newDatacenter.City != "Melbourne,VIC" {
			t.Errorf("expected new datacenter name Melbourne,VIC, received %s", newDatacenter.City)
		}
	})

	t.Run("Test failure on empty name for datacentre creation", func(t *testing.T) {

		creationResponse := `{"message":"datacenterName cannot be null","statusCode":400,"documentKind":"com:vmware:xenon:common:ServiceErrorResponse","errorCode":0}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusBadGateway,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(creationResponse))),
			}, nil
		})

		newDatacenter := types.Datacenter{Name: "", City: "Melbourne,VIC"}
		c := NewApiClient("127.0.0.1", true, cli)

		err := c.DatacenterService.Create(ctx, &newDatacenter)
		if err == nil {
			t.Errorf("expected error to be thrown when no name provided")
		}
	})

	t.Run("Test server response error for datacentre creation", func(t *testing.T) {

		creationResponse := `{"message":"datacenterName cannot be null","statusCode":400,"documentKind":"com:vmware:xenon:common:ServiceErrorResponse","errorCode":0}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusBadGateway,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(creationResponse))),
			}, nil
		})

		newDatacenter := types.Datacenter{Name: "TestDataCenter", City: "Melbourne,VIC"}
		c := NewApiClient("127.0.0.1", true, cli)

		err := c.DatacenterService.Create(ctx, &newDatacenter)
		if err == nil {
			t.Errorf("expected error to be thrown when no name provided")
		}
	})

	t.Run("Test encoding error for datacentre creation", func(t *testing.T) {

		//creationResponse := `{"message":"datacenterName cannot be null"}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
			}, nil
		})

		newDatacenter := types.Datacenter{Name: "TestDataCenter", City: "Melbourne,VIC"}
		c := NewApiClient("127.0.0.1", true, cli)

		err := c.DatacenterService.Create(ctx, &newDatacenter)
		if err == nil {
			t.Errorf("expected error to be thrown when no name provided")
		}
	})
}

func TestDatacenterAPIService_Delete(t *testing.T) {

	var ctx = context.Background()
	datacenterID := "2c551b23da979e75588bd59d7418b"
	expectedURL := "127.0.0.1/lcm/api/db/inventory/datacenter/" + datacenterID

	t.Run("Test successful datacenter deletion", func(t *testing.T) {

		serverResponse := `{"datacenterName":"rewdsaR","city":"Bangalore","state":"Karnataka","country":"IN","longitude":"77.59369","documentVersion":1,"documentKind":"com:vmware:vrealize:lcm:nxinventory:document:datacenter:DataCenter","documentSelfLink":"/lcm/api/db/inventory/datacenter/2c551b23da979e75588bd59d7418b","documentUpdateTimeMicros":1557812363707001,"documentUpdateAction":"DELETE","documentExpirationTimeMicros":1557812363707000,"documentAuthPrincipalLink":"/core/authz/users/vLCMAdmin"}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "DELETE" {
				return nil, fmt.Errorf("expected DELETE method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})


		c := NewApiClient("127.0.0.1", true, cli)

		err := c.DatacenterService.Delete(ctx, datacenterID)
		if err != nil {
			t.Errorf("expected no error when creating datacenter")
		}

	})

	t.Run("Test failure on empty name for datacentre deletion", func(t *testing.T) {

		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "DELETE" {
				return nil, fmt.Errorf("expected DELETE method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
			}, nil
		})


		c := NewApiClient("127.0.0.1", true, cli)

		err := c.DatacenterService.Delete(ctx, "")
		if err == nil {
			t.Errorf("expected error to be thrown when no name provided")
		}
	})

	t.Run("Test server response error for datacentre deletion", func(t *testing.T) {

		creationResponse := `{"message":"datacenterName cannot be null","statusCode":400,"documentKind":"com:vmware:xenon:common:ServiceErrorResponse","errorCode":0}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "DELETE" {
				return nil, fmt.Errorf("expected DELETE method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusBadGateway,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(creationResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.DatacenterService.Delete(ctx, datacenterID)
		if err == nil {
			t.Errorf("expected error to be thrown when no name provided")
		}
	})


}
