package datacenter

import (
	"encoding/json"
	"github.com/sdbrett/vrlcmsdk/types"
	"io/ioutil"
	"net/http"
)



func GetDatacenterResponse(r *http.Response) (types.Datacenter, error) {

	d := &types.Datacenter{}
	// Parse response body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return *d, err
	}

	// Marshall response body into loginResponse struct

	err = json.Unmarshal(body, &d)
	if err != nil {
		return *d, err
	}

	return *d, nil

}



