package client

import (
	"bytes"
	"fmt"
	"testing"
)

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
