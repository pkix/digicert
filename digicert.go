// Package digicert implements the DigiCert v2 API.
package digicert

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseURI = "https://www.digicert.com/services/v2/"

// Client as standard client
type Client struct {
	client     *http.Client
	statusCode int
	request    interface{}
	result     interface{}
	AuthKey    string
	headers    http.Header
}

// SchemeValidationErrors provides basic fields for scheme validation errors, the general error handling by return http status codes.
// Order Management - Error Codes & Responses https://www.digicert.com/services/v2/documentation/order/order-management-error-codes-and-messages
// Submitting Orders - Error Codes & Responses https://www.digicert.com/services/v2/documentation/order/submitting-orders-error-codes-and-messages
// Errors and Troubleshooting https://www.digicert.com/services/v2/documentation/errors
type SchemeValidationErrors struct {
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		SchemeValidationErrors
	}
}

// New exports digicert new api instance.
func New(key string) (*Client, error) {
	if key == "" {
		return nil, errors.New("The digicert api credentials must input")
	}
	c := &Client{
		AuthKey: key,
	}
	return c, nil

}

// apiconnect exports a http client dial to digicert api endpoints.
func (c *Client) makeRequest(method, uri string, headers http.Header) ([]byte, error) {
	var req *http.Request
	var err error
	fullURI := baseURI + strings.Trim(uri, "/")
	// log.Println("fullURI - ", fullURI)
	if method == "GET" || method == "DELETE" {
		req, err = http.NewRequest(method, fullURI, nil)
		if err != nil {
			return nil, err
		}
	} else {
		data, err := json.Marshal(c.request)
		if err != nil {
			return nil, err
		}
		streaming := bytes.NewBuffer([]byte(data))
		req, err = http.NewRequest(method, fullURI, streaming)
		if err != nil {
			return nil, err
		}
	}
	combinedHeaders := make(http.Header)
	copyHeader(combinedHeaders, c.headers)
	copyHeader(combinedHeaders, headers)
	req.Header = combinedHeaders
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("X-DC-DEVKEY", c.AuthKey)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
			InsecureSkipVerify:       false,
		},
	}
	c.client = &http.Client{Transport: tr}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	c.statusCode = res.StatusCode

	switch {
	case res.StatusCode == 401:
		return nil, errors.New("Unauthorized: Returned if the page is accessed without a valid API Key")
	case res.StatusCode == 403:
		return nil, errors.New("User doesn't have permission to perform the requested action")
	case res.StatusCode == 404:
		return nil, errors.New("Returned if the page doesn't exist or the API doesn't have permission to interact with a particular item")
	case res.StatusCode == 406:
		return nil, errors.New("If the client doesn't specify a valid acceptable content-type")
	case res.StatusCode == 429:
		return nil, errors.New("Too many requests. The client has sent too many requests in a given amount of time")
	case res.StatusCode == 500:
		return nil, errors.New("Unexpected behavior that the API couldn't recover from")
	case res.StatusCode == 503:
		return nil, errors.New("The system is currently unavailable")
	case res.StatusCode == 204:
		return nil, nil
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, err
}

// copyHeader copies all headers for `source` and sets them on `target`.
// based on https://godoc.org/github.com/golang/gddo/httputil/header#Copy
func copyHeader(target, source http.Header) {
	for k, vs := range source {
		target[k] = vs
	}
}
