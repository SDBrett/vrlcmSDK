package vrlcmSdk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSdkConnection_Login(t *testing.T) {
	authToken := "TOKENCODE"

	responseBody := `{"Token": "` + authToken + `"}`
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

	err := c.Login("username", "password")
	if err != nil {
		t.Errorf("no error response")
	}
	if c.Token != authToken {
		t.Errorf("expected token code %s, received %s", authToken, c.Token)
	}
}

func TestSdkConnection_Login2(t *testing.T) {

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/login" {
				http.Error(w, "Error", http.StatusUnauthorized)
			}

		}),
	)
	defer ts.Close()

	var c SdkConnection
	c = SdkConnection{BaseUrl: ts.URL, IgnoreCertError: true}

	err := c.Login("username", "password")

	if err == nil {
		t.Errorf("expected error response, no response received")
	}
}

func TestSdkConnection_Logout(t *testing.T) {

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/logout" {
				w.Header().Add("Content-Type", "application/json")
			}
		}),
	)
	defer ts.Close()

	var c SdkConnection
	c = SdkConnection{BaseUrl: ts.URL, IgnoreCertError: true}

	err := c.Logout()
	if err != nil {
		t.Errorf("no error response")
	}
}

func TestSdkConnection_Logout2(t *testing.T) {

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/logout" {
				http.Error(w, "Error", http.StatusUnauthorized)
			}

		}),
	)
	defer ts.Close()

	var c SdkConnection
	c = SdkConnection{BaseUrl: ts.URL, IgnoreCertError: true}

	err := c.Logout()

	if err == nil {
		t.Errorf("expected error response, no response received")
	}
}
