package vrlcmsdk

import (
	"net/http"
	"testing"
)

func TestNewDefaultApiTransport(t *testing.T) {

	testCases := []bool{true, false}

	for _, i := range testCases {

		rt := NewDefaultApiTransport(i)
		tr, err := rt.(*http.Transport)

		if !err {
			t.Errorf("got a %T, want an *http.Transport", rt)
		}

		if tr.TLSClientConfig.InsecureSkipVerify != i {
			t.Errorf("Excepted InsecureSkipVerify to equal %t, received %t", i, tr.TLSClientConfig.InsecureSkipVerify)
		}
	}
}

func TestApiClient_newDefaultClient(t *testing.T) {

	c := NewApiClient("https://192.168.17.128", true, nil)

	if c.httpClient == nil {
		t.Errorf("Api client contains no httpClient")
	}

}
