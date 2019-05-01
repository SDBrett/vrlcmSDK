package vrlcmsdk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func newHttpResponseCodes(count int) []int {

	rand.Seed(time.Now().UTC().UnixNano())
	var s []int
	min := 200
	max := 500

	for i := 0; i <= count; i++ {
		s = append(s, rand.Intn(max-min)+min)
		fmt.Println()
	}

	return s
}

func TestValidateHttpResponse(t *testing.T) {

	responseCodes := newHttpResponseCodes(20)
	r := ioutil.NopCloser(bytes.NewReader([]byte("hello world")))
	resp := &http.Response{Body: r}
	for _, c := range responseCodes {

		resp.StatusCode = c

		err := ValidateHttpResponse(*resp)

		if c <= 299 {
			if err != nil {
				t.Errorf("expected no error for statuscode : %d", resp.StatusCode)
			}
		}
		if c >= 300 {
			if err == nil {
				t.Errorf("expected error for status code : %d", resp.StatusCode)
			}
		}
	}

}

func TestApiClient_setDefaultHeaders(t *testing.T) {

	var c ApiClient
	c.setDefaultHeaders()

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
	c.token = mockToken
	c.addAuthHeader()

	if c.headers.Get("x-xenon-auth-token") != mockToken {
		t.Errorf("Expected x-xenon-auth-token header to equal %s received %s", mockToken, c.headers.Get("x-xenon-auth-token"))
	}

}
