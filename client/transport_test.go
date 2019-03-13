package client

import (
	"net/http"
	"testing"
)

func TestNewDefaultSdkTransport(t *testing.T) {

	testCases := []bool{true, false}

	for _, i := range testCases {

		rt := NewDefaultSdkTransport(i)
		tr, ok := rt.(*http.Transport)

		if !ok {
			t.Fatalf("got a %T, want an *http.Transport", rt)
		}
		if tr.TLSClientConfig.InsecureSkipVerify != i {
			t.Fatal("oops")
		}
	}

}
