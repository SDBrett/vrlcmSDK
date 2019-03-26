package vrlcmSdk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSdkConnection_GetDatacenters(t *testing.T) {


	responseBody := `[{"id": "12345", "name": "test"}]`
	
	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/login" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(responseBody))
			}

		}),
	)
	defer ts.Close()

	var c SdkConnection
	c = SdkConnection{BaseUrl: ts.URL, IgnoreCertError: true}
	c.Token = "test"
	c.newSdkHeaders()
	d, err := c.GetDatacenters()
	
	
	if err != nil {
		t.Errorf("no error response")
	}
	if d.Datacenter == nil {
		t.Errorf("expected token code %s, received %s", "authToken", c.Token)
	}
}

func TestGet()  {
	
}