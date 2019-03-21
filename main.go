package main

import (
	"fmt"
	"vrlcmSDK/vrlcmSdk"
)

func main() {

	var c vrlcmSdk.SdkConnection
	c = vrlcmSdk.SdkConnection{BaseUrl: "https://192.168.17.145/lcm/api/v1", IgnoreCertError: true}

	err := c.Login("admin@localhost", "vmware")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

	c.Logout()

}
