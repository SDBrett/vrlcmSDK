package vrlcmsdk

import (
	"context"
	"encoding/json"
	"github.com/sdbrett/vrlcmsdk/types"
)

// Vcenter API service
type SettingsAPIService service

type SshPassword struct {
	SshUserPassword string `json:"sshuserPassword"`
}

type rootPassword struct {
	RootPassword string `json:"rootPassword"`
}

// Set new root password
func (service *SettingsAPIService) SetRootPassword(ctx context.Context, password string) error {

	url := service.client.basePath + "/settings"
	body := rootPassword{RootPassword: password}

	resp, err := service.client.post(ctx, url, body, *service.client.headers)
	if err != nil {
		return err
	}

	ensureReaderClosed(resp)

	return nil

}

// Set new admin password
func (service *SettingsAPIService) SetAdminPassword(ctx context.Context, password string) error {

	url := service.client.basePath + "/settings"
	body := rootPassword{RootPassword: password}

	resp, err := service.client.post(ctx, url, body, *service.client.headers)
	if err != nil {
		return err
	}

	ensureReaderClosed(resp)

	return nil

}

// Get LCM configuration settings
func (service *SettingsAPIService) GetLcmConfigSettings(ctx context.Context) (types.ConfigSettings, error) {

	url := service.client.Host + "/lcm/api/settings/lcmconfig"
	settings := types.ConfigSettings{}
	resp, err := service.client.get(ctx, url, *service.client.headers)
	if err != nil {
		return settings, err
	}

	err = json.NewDecoder(resp.body).Decode(&settings)
	if err != nil {
		return settings, err
	}
	ensureReaderClosed(resp)

	return settings, nil

}

// Get LCM network configuration settings
func (service *SettingsAPIService) GetNetworkStatus(ctx context.Context) (types.NetworkStatus, error) {

	url := service.client.Host + "/lcm/gui/api/lcmVaNetworkStatus"
	settings := types.NetworkStatus{}
	resp, err := service.client.get(ctx, url, *service.client.headers)
	if err != nil {
		return settings, err
	}

	err = json.NewDecoder(resp.body).Decode(&settings)
	if err != nil {
		return settings, err
	}
	ensureReaderClosed(resp)

	return settings, nil

}
