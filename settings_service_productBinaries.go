package vrlcmsdk

import (
	"context"
	"encoding/json"
	"github.com/sdbrett/vrlcmsdk/types"
)

// Trigger refresh patch request
func (service *SettingsAPIService) RefreshPatches(ctx context.Context) (types.Request, error) {

	url := service.client.Host + "/lcm/gui/api/action/refreshPatches"
	var request types.Request

	resp, err := service.client.post(ctx, url, nil, *service.client.headers)
	if err != nil {
		return request, err
	}

	err = json.NewDecoder(resp.body).Decode(&request)
	if err != nil {
		return request, err
	}

	ensureReaderClosed(resp)

	return request, nil

}
