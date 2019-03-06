package client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestValidateHttpResponse(t *testing.T) {

	r := ioutil.NopCloser(bytes.NewReader([]byte("hello world")))
	resp := &http.Response{StatusCode: 400, Body: r}
	err := ValidateHttpResponse(*resp)
	if err == nil {
		t.Errorf("Expected error response for status: %d", resp.StatusCode)
	}

	resp = &http.Response{StatusCode: 200, Body: r}
	err = ValidateHttpResponse(*resp)
	if err != nil {
		t.Errorf("Expected noerror response for status: %d", resp.StatusCode)
	}
}

func TestCreateLoginRequestBody(t *testing.T) {

	username := "admin@localhost"
	password := "password"
	bodyString := `{"username":"` + username + `", "password":"` + password + `"}`

	testBody := []byte(bodyString)

	functionReturn := CreateLoginRequestBody(username, password)
	eql := bytes.Equal(testBody, functionReturn)
	fmt.Println(eql)

	if !bytes.Equal(testBody, functionReturn) {
		t.Errorf("Error comparing")
	}

}

func TestNewDefaultSdkTransport(t *testing.T) {

	testValues := []bool{true, false}

	for _, i := range testValues {
		transport := NewDefaultSdkTransport(i)

		if transport.TLSClientConfig.InsecureSkipVerify != i {
			t.Errorf("InsecureVerify set to: %t, expected %t", transport.TLSClientConfig.InsecureSkipVerify, i)
		}
	}
}

func TestNewDefaultSdkClient(t *testing.T) {

	trans := MockNewDefaultSdkTransport(false)

	c := NewDefaultSdkClient(&trans)

	if c.Transport == nil {
		t.Errorf("Expected transport not to be nil")

	}
}

func MockNewDefaultSdkTransport(skipCertVerify bool) http.Transport {
	t := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 90 * time.Second,
			DualStack: true,
		}).DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipCertVerify},
	}

	return *t
}
