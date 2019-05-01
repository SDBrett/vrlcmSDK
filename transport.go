package vrlcmsdk

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

//Default Http transport configuration
func NewDefaultApiTransport(skipCertVerify bool) http.RoundTripper {
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



func (c *ApiClient) newDefaultClient() {

	c.httpClient = &http.Client{Transport: NewDefaultApiTransport(c.AllowInsecure)}

}
