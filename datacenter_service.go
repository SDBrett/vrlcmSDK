package vrlcmsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sdbrett/vrlcmsdk/types"
	"log"
)

type DatacenterAPIService service

func (service *DatacenterAPIService) GetAllDatacenters(ctx context.Context) (types.Datacenters, error) {

	url := service.client.basePath + "/view/datacenter"
	d := types.Datacenters{}

	resp, err := service.client.get(ctx, url, *service.client.headers)

	if err != nil {
		return d, err
	}

	err = json.NewDecoder(resp.body).Decode(&d.Datacenter)
	ensureReaderClosed(resp)

	for x, i := range d.Datacenter {
		id := i.ID
		d.Datacenter[x], err = service.GetDatacenter(ctx, id)
		if err != nil {
			log.Printf("received error getting datacenter for datacenter id %s", id)
			return d, err
		}
		d.Datacenter[x].ID = id
		fmt.Printf(id)
	}
	return d, nil

}

func (service *DatacenterAPIService) GetDatacenter(ctx context.Context, id string) (types.Datacenter, error) {

	url := service.client.basePath + "/view/datacenter?datacenterId=" + id
	d := types.Datacenter{}

	resp, err := service.client.get(ctx, url, *service.client.headers)
	if err != nil {
		return d, err
	}

	err = json.NewDecoder(resp.body).Decode(&d)
	if err != nil {
		return d, err
	}

	return d, nil

}

func (service *DatacenterAPIService) Create(ctx context.Context, d *types.Datacenter) error {

	url := service.client.basePath + "/action/create/datacenter"
	var tempDC types.Datacenter
	if d.Name == "" {
		err := errors.New("Datacenter name cannot be empty")
		return err
	}

	resp, err := service.client.post(ctx, url, d, *service.client.headers)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.body).Decode(&tempDC)
	if err != nil {
		return err
	}

	d.ID = tempDC.ID
	ensureReaderClosed(resp)

	return nil

}
