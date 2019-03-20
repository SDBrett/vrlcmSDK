package vrlcmSdk

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

//Default Http transport configuration
func NewDefaultSdkTransport(skipCertVerify bool) http.RoundTripper {
	t := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 90 * time.Second,
			DualStack: true,
		}).DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipCertVerify},
	}

	return t
}

func NewApiClient(options ...func(*ApiClient)) (*ApiClient, error) {
	c := http.Client{}

	client := ApiClient{Client: c}

	for _, option := range options {
		option(&client)
	}

	return &client, nil
}
