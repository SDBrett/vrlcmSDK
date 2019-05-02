package vrlcmsdk

import "testing"

func TestNewApiClient(t *testing.T) {

	ignoreSSL := true
	host := "https://192.168.17.128"

	c := NewApiClient(host, ignoreSSL, nil)

	if c.AllowInsecure != ignoreSSL {
		t.Errorf("SSL verify setting does not match input, expected %t, received %t", c.AllowInsecure, ignoreSSL)
	}

	if c.Host != host {
		t.Errorf("Host setting does not match input, expected %s, received %s", c.Host, host)
	}

}
