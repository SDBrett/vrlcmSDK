package vrlcmsdk

import (
	"net/http"
	"testing"
)

func TestNewApiClient(t *testing.T) {

	t.Run("Test create new client without http client", func(t *testing.T) {

		ignoreSSL := true
		host := "https://192.168.17.128"

		c := NewApiClient(host, ignoreSSL, nil)

		if c.AllowInsecure != ignoreSSL {
			t.Errorf("SSL verify setting does not match input, expected %t, received %t", c.AllowInsecure, ignoreSSL)
		}

		if c.Host != host {
			t.Errorf("Host setting does not match input, expected %s, received %s", c.Host, host)
		}
	})

	t.Run("Test create new client with http client", func(t *testing.T) {

		ignoreSSL := true
		host := "https://192.168.17.128"

		client := &http.Client{}

		c := NewApiClient(host, ignoreSSL, client)

		if c.Host != host {
			t.Errorf("Host setting does not match input, expected %s, received %s", c.Host, host)
		}
	})

}
