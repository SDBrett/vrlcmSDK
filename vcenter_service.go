package vrlcmsdk

import (
	"context"
	"encoding/json"
	"github.com/sdbrett/vrlcmsdk/types"
)

// Vcenter API service
type VcenterAPIService service

// Create Vcenter server object, returns request object
func (service *VcenterAPIService) Create(ctx context.Context, vc *types.Vcenter) (types.Request, error) {

	url := service.client.basePath + "/action/add/vc"
	var request types.Request


	resp, err := service.client.post(ctx, url, vc, *service.client.headers)
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
