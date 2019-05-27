package vrlcmsdk

import (
	"context"
	"encoding/json"
	"github.com/sdbrett/vrlcmsdk/types"
)

// LCM Settings API service
type SettingsAPIService service

// Set new root password
func (service *SettingsAPIService) SetRootPassword(ctx context.Context, password string) error {

	url := service.client.basePath + "/v1/settings"
	body := types.RootPassword{RootPassword: password}

	resp, err := service.client.post(ctx, url, body, *service.client.headers)
	if err != nil {
		return err
	}

	ensureReaderClosed(resp)

	return nil

}

// Set new admin password
func (service *SettingsAPIService) SetAdminPassword(ctx context.Context, password string) error {

	url := service.client.basePath + "/v1/settings"
	body := types.RootPassword{RootPassword: password}

	resp, err := service.client.post(ctx, url, body, *service.client.headers)
	if err != nil {
		return err
	}

	ensureReaderClosed(resp)

	return nil

}

// Get LCM configuration settings
func (service *SettingsAPIService) GetLcmConfigSettings(ctx context.Context) (types.ConfigSettings, error) {

	url := service.client.basePath + "/settings/lcmconfig"
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

// Set new admin password
func (service *SettingsAPIService) SetRestartSchedule(ctx context.Context, schedule types.RestartSchedule) error {

	url := service.client.basePath + "/maintenance/xserver-restart-config"

	resp, err := service.client.post(ctx, url, schedule, *service.client.headers)
	if err != nil {
		return err
	}

	ensureReaderClosed(resp)

	return nil

}