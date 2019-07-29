package vrlcmsdk

import (
	"testing"
)

func TestApiClient_setDefaultHeaders(t *testing.T) {

	var c ApiClient
	c.SetDefaultHeaders()

	if c.headers.Get("Accept") != "application/json" {
		t.Errorf("Expected Accept header to equal \"application\\json\" received %s", c.headers.Get("Accept"))
	}
	if c.headers.Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type header to equal \"application\\json\" received %s", c.headers.Get("Content-Type"))
	}
}

func TestApiClient_addAuthHeader(t *testing.T) {

	mockToken := "MOCKTOKEN"

	c := NewApiClient("https://192.168.17.128", true, nil)
	c.Token = mockToken
	c.AddAuthHeader()

	if c.headers.Get("x-xenon-auth-Token") != mockToken {
		t.Errorf("Expected x-xenon-auth-Token header to equal %s received %s", mockToken, c.headers.Get("x-xenon-auth-Token"))
	}

}
