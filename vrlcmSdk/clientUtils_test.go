package vrlcmSdk

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

func TestSdkConnection_newSdkHeaders(t *testing.T) {
	mockToken := "TOKENTEST"
	mockContentType := "application/json"

	var c SdkConnection
	c = SdkConnection{Token: mockToken}

	c.newSdkHeaders()

	if c.headers.Get("x-xenon-auth-token") != mockToken {
		t.Errorf("expected token code: %s. received %s", mockToken, c.headers.Get("x-xenon-auth-token"))
	}
	if c.headers.Get("Content-Type") != mockContentType {
		t.Errorf("expected token code: %s. received %s", mockContentType, c.headers.Get("Content-Type"))
	}
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
