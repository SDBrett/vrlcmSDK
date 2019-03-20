package vrlcmSdk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	okResponse = "TOKENCODE"
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

func TestDoStuffWithRoundTripper(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := `{"token":"TOKENCODE"}`
		_, err := fmt.Fprintln(w, body)
		if err != nil {
			t.Fatal(err)
		}
	}))

	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	l := &SdkConnection{BaseUrl: "https://192.168.17.145/lcm/api/v1", IgnoreCertError: true}

	err = loginResponse(res, l)
	if err != nil {
		t.Fatal(err)
	}

	if l.Token != okResponse {
		t.Fatalf("Expected response of %s but received %s", okResponse, l.Token)
	}
}
