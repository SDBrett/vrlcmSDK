package vrlcmSdk

import (
	"bytes"
	"github.com/hashicorp/packer/common/json"
	"io/ioutil"
	"net/http"
)

type SdkConnection struct {
	BaseUrl         string
	Token           string
	IgnoreCertError bool
	Client http.Client
	headers	http.Header
}


type LoginResponse struct {
	Token string `json:"token"`
}

func (s *SdkConnection) Login(u, p string) error {

	t := func(c *SdkConnection) {
		c.Client.Transport = NewDefaultSdkTransport(s.IgnoreCertError)
	}

	c, _ := NewApiClient(t)

	url := s.BaseUrl + "/login"

	body := CreateLoginRequestBody(u, p)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	response, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	err = ValidateHttpResponse(*response)
	if err != nil {
		return err
	}
	err = getLoginResponse(response, s)
	if err != nil {
		return err
	}

	s.newSdkHeaders()

	return nil
}



func getLoginResponse(r *http.Response, s *SdkConnection) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	loginResponse := LoginResponse{}

	err = json.Unmarshal(body, &loginResponse)

	if err != nil {
		return err
	}

	s.Token = loginResponse.Token

	return nil

}

func CreateLoginRequestBody(u, p string) []byte {
	bodyString := `{"username":"` + u + `", "password":"` + p + `"}`
	return []byte(bodyString)
}

/*
// TODO Add result parsing and return statement
func (s *SdkConnection) GetDataCenters() error {

	t := func(c *SdkConnection) {
		c.Client.Transport = NewDefaultSdkTransport(s.IgnoreCertError)
	}

	c, _ := NewApiClient(t)

	url := s.BaseUrl + "/view/datacenter"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	response, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	err = ValidateHttpResponse(*response)
	if err != nil {
		return err
	}

}

*/


func (s *SdkConnection) Logout() error {

	t := func(c *SdkConnection) {
		c.Client.Transport = NewDefaultSdkTransport(s.IgnoreCertError)
	}

	c, _ := NewApiClient(t)

	url := s.BaseUrl + "/logout"


	req, err := http.NewRequest("POST", url, nil )
	if err != nil {
		return err
	}
	req.Header = s.headers

	response, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	err = ValidateHttpResponse(*response)
	if err != nil {
		return err
	}

	return nil
}



