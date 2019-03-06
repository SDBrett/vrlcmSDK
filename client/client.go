package client

import (
	"bytes"
	"github.com/hashicorp/packer/common/json"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Login(u, p string) error
}

type SdkClient struct {
	BaseUrl         string
	Token           string
	IgnoreCertError bool
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (s *SdkClient) Login(u, p string) error {

	t := NewDefaultSdkTransport(s.IgnoreCertError)
	c := NewDefaultSdkClient(t)
	url := s.BaseUrl + "/login"

	body := CreateLoginRequestBody(u, p)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	response, err := c.Do(req)
	if err != nil {
		return err
	}

	err = ValidateHttpResponse(*response)
	if err != nil {
		return err
	}
	err = loginResponse(response, s)
	if err != nil {
		return err
	}

	return nil
}

func loginResponse(r *http.Response, s *SdkClient) error {

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
func (c *Client) GetDataCenter() error {
	url := c.baseURL + "/view/datacenter"


	resp, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}

	defer Close(resp.Body)

	return nil
}

*/
