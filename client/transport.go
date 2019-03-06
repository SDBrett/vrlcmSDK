package client

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

//Default Http transport configuration
func NewDefaultSdkTransport(skipCertVerify bool) http.Transport {
	t := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 90 * time.Second,
			DualStack: true,
		}).DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipCertVerify},
	}

	return *t
}

// Default Http client configuration
func NewDefaultSdkClient(t http.Transport) http.Client {
	c := &http.Client{
		Transport: &t,
	}

	return *c
}
