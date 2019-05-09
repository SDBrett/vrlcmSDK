package vrlcmsdk

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateLoginRequestBody(t *testing.T) {

}
func TestApiClient_Login(t *testing.T) {

	authToken := "TOKENCODE"

	responseBody := `{"Token": "` + authToken + `"}`
	// Mock http server

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/lcm/api/v1/login" {
				w.Header().Add("Content-Type", "application/json")
				_, err := w.Write([]byte(responseBody))
				if err != nil {
					log.Fatalf("Error writing response")
				}
			}
			if r.URL.Path == "/fail/lcm/api/v1/login" {
				w.Header().Add("Content-Type", "application/json")
				http.Error(w, "Error", http.StatusUnauthorized)
			}
			if r.URL.Path == "/404/lcm/api/v1/login" {
				http.Error(w, "Error", http.StatusNotFound)
			}

		}),
	)
	defer ts.Close()

	t.Run("Test successful login", func(t *testing.T) {

		ignoreSSL := true
		c := NewApiClient(ts.URL, ignoreSSL, nil)

		err := c.Login(ctx, "admin@localhost", "vmware")
		if err != nil {
			t.Errorf("no error response")
		}
		if c.token != authToken {
			t.Errorf("expected token code %s, received %s", authToken, c.token)
		}
	})

	t.Run("Test successful login", func(t *testing.T) {

		ignoreSSL := true
		uri := ts.URL + "/fail"

		c := NewApiClient(uri, ignoreSSL, nil)

		err := c.Login(ctx, "admin@localhost", "vmware")
		if err == nil {
			t.Errorf("expected error response, no response received")
		}
	})

	t.Run("Test bad url login", func(t *testing.T) {

		ignoreSSL := true
		uri := ts.URL + "/404"
		c := NewApiClient(uri, ignoreSSL, nil)

		err := c.Login(ctx, "username", "password")
		if err == nil {
			t.Errorf("expected error response, no response received")
		}
	})

}

func TestApiClient_Logout(t *testing.T) {

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/lcm/api/v1/login/logout" {
				w.Header().Add("Content-Type", "application/json")
			}
			if r.URL.Path == "/fail/lcm/api/v1/logout" {
				http.Error(w, "Error", http.StatusUnauthorized)
			}
			if r.URL.Path == "/404/lcm/api/v1/logout" {
				http.Error(w, "Error", http.StatusNotFound)
			}

		}),
	)
	defer ts.Close()

	t.Run("Test successful logout", func(t *testing.T) {

		ignoreSSL := true
		c := NewApiClient(ts.URL, ignoreSSL, nil)
		err := c.Logout(ctx)
		if err != nil {
			t.Errorf("no error response")
		}
	})
	t.Run("Test failed logout", func(t *testing.T) {

		uri := ts.URL + "/fail"
		ignoreSSL := true

		c := NewApiClient(uri, ignoreSSL, nil)
		err := c.Logout(ctx)
		if err == nil {
			t.Errorf("expected error response, no response received")
		}
	})

	t.Run("Test bad URI logout", func(t *testing.T) {

		uri := "https://fail"
		ignoreSSL := true

		c := NewApiClient(uri, ignoreSSL, nil)
		err := c.Logout(ctx)
		if err == nil {
			t.Errorf("expected error response, no response received")
		}
	})

	t.Run("Test bad url logout", func(t *testing.T) {

		uri := ts.URL + "/404"
		ignoreSSL := true

		c := NewApiClient(uri, ignoreSSL, nil)
		err := c.Logout(ctx)
		if err == nil {
			t.Errorf("expected error response, no response received")
		}
	})
}
