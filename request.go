// Large portions of this code originated from the Docker Moby project

package vrlcmsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sdbrett/vrlcmsdk/types"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// serverResponse is a wrapper for http API responses.
type serverResponse struct {
	body       io.ReadCloser
	header     http.Header
	statusCode int
	reqURL     *url.URL
}

// head sends an http request to the docker API using the method HEAD.
func (cli *ApiClient) head(ctx context.Context, url string, headers http.Header) (serverResponse, error) {
	return cli.sendRequest(ctx, "HEAD", url, nil, headers)

}

// get sends an http request to the docker API using the method GET with a specific Go context.
func (cli *ApiClient) get(ctx context.Context, url string, headers http.Header) (serverResponse, error) {
	return cli.sendRequest(ctx, "GET", url, nil, headers)
}

// post sends an http request to the docker API using the method POST with a specific Go context.
func (cli *ApiClient) post(ctx context.Context, url string, obj interface{}, headers http.Header) (serverResponse, error) {
	body, headers, err := encodeBody(obj, headers)
	if err != nil {
		return serverResponse{}, err
	}
	return cli.sendRequest(ctx, "POST", url, body, headers)
}

func (cli *ApiClient) postRaw(ctx context.Context, url string, body io.Reader, headers http.Header) (serverResponse, error) {
	return cli.sendRequest(ctx, "POST", url, body, headers)
}

// put sends an http request to the docker API using the method PUT.
func (cli *ApiClient) put(ctx context.Context, url string, obj interface{}, headers http.Header) (serverResponse, error) {
	body, headers, err := encodeBody(obj, headers)
	if err != nil {
		return serverResponse{}, err
	}
	return cli.sendRequest(ctx, "PUT", url, body, headers)
}

// putRaw sends an http request to the docker API using the method PUT.
func (cli *ApiClient) putRaw(ctx context.Context, url string, body io.Reader, headers http.Header) (serverResponse, error) {
	return cli.sendRequest(ctx, "PUT", url, body, headers)
}

// delete sends an http request to the docker API using the method DELETE.
func (cli *ApiClient) delete(ctx context.Context, url string, headers http.Header) (serverResponse, error) {
	return cli.sendRequest(ctx, "DELETE", url, nil, headers)
}

func encodeBody(obj interface{}, headers http.Header) (io.Reader, http.Header, error) {
	if obj == nil {
		return nil, headers, nil
	}

	body, err := encodeData(obj)
	if err != nil {
		return nil, headers, err
	}
	if headers == nil {
		headers = make(map[string][]string)
	}
	headers["Content-Type"] = []string{"application/json"}
	return body, headers, nil
}

func (cli *ApiClient) buildRequest(method, url string, body io.Reader, headers http.Header) (*http.Request, error) {
	expectedPayload := method == "POST" || method == "PUT"
	if expectedPayload && body == nil {
		body = bytes.NewReader([]byte{})
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = headers

	return req, nil
}

func (cli *ApiClient) sendRequest(ctx context.Context, method, url string, body io.Reader, headers http.Header) (serverResponse, error) {
	req, err := cli.buildRequest(method, url, body, headers)
	if err != nil {
		return serverResponse{}, err
	}

	resp, err := cli.doRequest(ctx, req)
	if err != nil {
		errorMessage := errors.New("Error making request" + err.Error())
		return serverResponse{}, errorMessage

	}

	err = cli.checkResponseErr(resp)

	return resp, err

}

func (cli *ApiClient) doRequest(ctx context.Context, req *http.Request) (serverResponse, error) {
	serverResp := serverResponse{statusCode: -1, reqURL: req.URL}

	req = req.WithContext(ctx)
	resp, err := cli.httpClient.Do(req)
	if err != nil {
		return serverResp, err
	}

	if resp != nil {
		serverResp.statusCode = resp.StatusCode
		serverResp.body = resp.Body
		serverResp.header = resp.Header
	}
	return serverResp, nil
}

func (cli *ApiClient) checkResponseErr(serverResp serverResponse) error {
	if serverResp.statusCode >= 200 && serverResp.statusCode < 400 {
		return nil
	}

	var body []byte
	var err error
	if serverResp.body != nil {
		bodyMax := 1 * 1024 * 1024 // 1 MiB
		bodyR := &io.LimitedReader{
			R: serverResp.body,
			N: int64(bodyMax),
		}
		body, err = ioutil.ReadAll(bodyR)
		if err != nil {
			return err
		}
		if bodyR.N == 0 {
			return fmt.Errorf("request returned %s with a message (> %d bytes) for API route and version %s, check if the server supports the requested API version", http.StatusText(serverResp.statusCode), bodyMax, serverResp.reqURL)
		}
	}
	if len(body) == 0 {
		return fmt.Errorf("request returned %s for API route and version %s, check if the server supports the requested API version", http.StatusText(serverResp.statusCode), serverResp.reqURL)
	}

	var errorMessage string
	var errorResponse types.ErrorResponse
	if err := json.Unmarshal(body, &errorResponse); err != nil {
		return errors.Wrap(err, "Error reading JSON")
	}

	return errors.Wrap(errors.New(errorMessage), "Error response from server")
}

func encodeData(data interface{}) (*bytes.Buffer, error) {
	params := bytes.NewBuffer(nil)
	if data != nil {
		if err := json.NewEncoder(params).Encode(data); err != nil {
			return nil, err
		}
	}
	return params, nil
}

func ensureReaderClosed(response serverResponse) {
	if response.body != nil {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, response.body, 512)
		response.body.Close()
	}
}
