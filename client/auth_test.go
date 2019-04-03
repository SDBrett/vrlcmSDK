package client

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAuthToken(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	_, err := getAuthToken(resp)
	if err == nil {
		t.Errorf("Expected error")
	}

}

func TestSdkConnection_Login(t *testing.T) {
	authToken := "TOKENCODE"

	responseBody := `{"Token": "` + authToken + `"}`
	// Mock http server

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/login" {
				w.Header().Add("Content-Type", "application/json")
				_, err := w.Write([]byte(responseBody))
				if err != nil {
					log.Fatalf("Error writing response")
				}
			}
			if r.URL.Path == "/fail/login" {
				w.Header().Add("Content-Type", "application/json")
				http.Error(w, "Error", http.StatusUnauthorized)
			}
		}),
	)
	defer ts.Close()

	var c SdkConnection
	t.Run("Test successful login", func(t *testing.T) {
		c = SdkConnection{BaseUrl: ts.URL, IgnoreCertError: true}

		err := c.Login("username", "password")
		if err != nil {
			t.Errorf("no error response")
		}
		if c.Token != authToken {
			t.Errorf("expected token code %s, received %s", authToken, c.Token)
		}
	})

	t.Run("Test successful login", func(t *testing.T) {
		uri := ts.URL + "/fail"
		c = SdkConnection{BaseUrl: uri, IgnoreCertError: true}

		err := c.Login("username", "password")

		if err == nil {
			t.Errorf("expected error response, no response received")
		}
	})
	t.Run("Test bad url login", func(t *testing.T) {
		uri := ts.URL + "/fai @#$#%&^^&$@!@$^&% l"
		c = SdkConnection{BaseUrl: uri, IgnoreCertError: true}

		err := c.Login("username", "password")

		if err == nil {
			t.Errorf("expected error response, no response received")
		}
	})
}

func TestSdkConnection_Logout(t *testing.T) {

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/logout" {
				w.Header().Add("Content-Type", "application/json")
			}
			if r.URL.Path == "/fail/logout" {
				http.Error(w, "Error", http.StatusUnauthorized)
			}
		}),
	)
	defer ts.Close()

	var c SdkConnection
	t.Run("Test successful logout", func(t *testing.T) {
		c = SdkConnection{BaseUrl: ts.URL, IgnoreCertError: true}
		err := c.Logout()
		if err != nil {
			t.Errorf("no error response")
		}
	})
	t.Run("Test failed logout", func(t *testing.T) {
		uri := ts.URL + "/fail"
		c = SdkConnection{BaseUrl: uri, IgnoreCertError: true}
		err := c.Logout()
		if err == nil {
			t.Errorf("expected error response, no response received")
		}
	})
	t.Run("Test bad URI logout", func(t *testing.T) {
		uri := "https://fail"
		c = SdkConnection{BaseUrl: uri, IgnoreCertError: false}
		err := c.Logout()
		if err == nil {
			t.Errorf("expected error response, no response received")
		}
	})
	t.Run("Test bad url logout", func(t *testing.T) {
		uri := ts.URL + "/fai @#$#%&^^&$@!@$^&% l"
		c = SdkConnection{BaseUrl: uri, IgnoreCertError: true}

		err := c.Logout()
		if err == nil {
			t.Errorf("expected error response, no response received")
		}
	})
}
