package main

import (
	"fmt"
	"github.com/goharbor/harbor/src/common/utils/log"
	"vrlcmSDK/vrlcmSdk"
)

func main() {

	var c vrlcmSdk.SdkConnection
	c = vrlcmSdk.SdkConnection{BaseUrl: "https://192.168.17.145/lcm/api/v1", IgnoreCertError: true}

	err := c.Login("admin@localhost", "vmware")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(c)
	var d  vrlcmSdk.Datacenters
	err = d.GetDatacenters(c)
	if err != nil {
		log.Errorf("received error: %s", err)
	}
	fmt.Println(d.Datacenter[0].ID)

}
