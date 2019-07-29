package vrlcmsdk

import (
	"context"
	"encoding/json"
	"github.com/sdbrett/vrlcmsdk/types"
)

// Performs authentication function with vRLCM server
// Adds auth Token string to the ApiClient
func (cli *ApiClient) Login(ctx context.Context, u, p string) error {

	url := cli.basePath + "/v1/login"
	body := types.LoginBody{Username: u, Password: p}
	// body := CreateLoginRequestBody(u, p)

	resp, err := cli.post(ctx, url, body, *cli.headers)
	if err != nil {
		return err
	}

	loginToken := types.LoginResponse{}
	err = json.NewDecoder(resp.body).Decode(&loginToken)
	if err != nil {
		return err
	}
	cli.Token = loginToken.Token
	cli.AddAuthHeader()

	return nil
}

// Performs logout action against vRLCM server
func (cli *ApiClient) Logout(ctx context.Context) error {

	url := cli.basePath + "/v1/logout"

	_, err := cli.post(ctx, url, nil, *cli.headers)
	if err != nil {
		return err
	}

	return nil
}
