package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDatacenters_GetDatacenters(t *testing.T) {

	responseBody := `[{"id": "12345", "name": "test"}]`

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/view/datacenter" {
				w.Header().Add("Content-Type", "application/json")
				_, err := w.Write([]byte(responseBody))
				if err != nil {
					t.Errorf("error response: %s", err)
				}
			}
			if r.URL.Path == "/403/view/datacenter" {
				http.Error(w, "Error", http.StatusUnauthorized)
			}

		}),
	)
	defer ts.Close()

	t.Run("OK Process", func(t *testing.T) {
		var c SdkConnection
		c = SdkConnection{BaseUrl: ts.URL, IgnoreCertError: true}
		c.Token = "authToken"
		c.newSdkHeaders()
		c.newDefaultClient()

		var d Datacenters
		err := d.GetDatacenters(c)
		if err != nil {
			t.Errorf("error response: %s", err)
		}
	})
	t.Run("API Error", func(t *testing.T) {
		var c SdkConnection
		url := ts.URL + "/403"
		c = SdkConnection{BaseUrl: url, IgnoreCertError: true}
		c.Token = "authToken"
		c.newSdkHeaders()
		c.newDefaultClient()

		var d Datacenters
		err := d.GetDatacenters(c)
		if err == nil {
			t.Errorf("expected error, received nil")
		}
	})
}

func TestDatacenters_GetDatacenters2(t *testing.T) {

	responseBody := `{[{"id": "12345", "name": "test"}]}`

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/view/datacenter" {
				w.Header().Add("Content-Type", "application/json")
				_, err := w.Write([]byte(responseBody))
				if err != nil {
					t.Errorf("error response: %s", err)
				}
			}
		}),
	)
	defer ts.Close()

	t.Run("OK Process", func(t *testing.T) {
		var c SdkConnection
		c = SdkConnection{BaseUrl: ts.URL, IgnoreCertError: true}
		c.Token = "authToken"
		c.newSdkHeaders()
		c.newDefaultClient()

		var d Datacenters
		err := d.GetDatacenters(c)
		if err == nil {
			t.Errorf("expected error but received nil")
		}
	})

}
