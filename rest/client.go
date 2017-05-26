package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultEndpoint  = "https://my.zerotier.com/api/"
	defaultUserAgent = "zerotier-go"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	httpClient Doer

	Endpoint *url.URL

	APIToken string

	UserAgent string

	common service

	Network *NetworkService
	Self    *SelfService
}

// NewClient constructs and returns a reference to an instantiated Client.
func NewClient(httpClient Doer, options ...func(*Client)) *Client {
	endpoint, _ := url.Parse(defaultEndpoint)

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		httpClient: httpClient,
		Endpoint:   endpoint,
		UserAgent:  defaultUserAgent,
	}

	c.common.client = c
	c.Self = (*SelfService)(&c.common)
	c.Network = (*NetworkService)(&c.common)

	for _, option := range options {
		option(c)
	}
	return c
}

type service struct {
	client *Client
}

// SetHTTPClient sets a Client instances' httpClient.
func SetHTTPClient(httpClient Doer) func(*Client) {
	return func(c *Client) { c.httpClient = httpClient }
}

// SetAPIKey sets a Client instances' APIKey.
func SetAPIToken(token string) func(*Client) {
	return func(c *Client) { c.APIToken = token }
}

// SetEndpoint sets a Client instances' Endpoint.
func SetEndpoint(endpoint string) func(*Client) {
	return func(c *Client) { c.Endpoint, _ = url.Parse(endpoint) }
}

// SetUserAgent sets a Client instances' user agent.
func SetUserAgent(ua string) func(*Client) {
	return func(c *Client) { c.UserAgent = ua }
}

// Do satisfies the Doer interface.
func (c Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		// Try to unmarshal body into given type using streaming decoder.
		if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
			return nil, err
		}
	}

	return resp, err
}

// NewRequest constructs and returns a http.Request.
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	uri := c.Endpoint.ResolveReference(rel)

	// Encode body as json
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, uri.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "bearer "+c.APIToken)
	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// Response wraps stdlib http response.
type Response struct {
	*http.Response
}

// Error contains all http responses outside the 2xx range.
type Error struct {
	Resp    *http.Response
	Message string
}

// Satisfy std lib error interface.
func (re *Error) Error() string {
	return fmt.Sprintf("%v %v: %d %v", re.Resp.Request.Method, re.Resp.Request.URL, re.Resp.StatusCode, re.Message)
}

// CheckResponse handles parsing of rest api errors. Returns nil if no error.
func CheckResponse(resp *http.Response) error {
	if c := resp.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	restErr := &Error{Resp: resp}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(b) == 0 {
		return restErr
	}

	err = json.Unmarshal(b, restErr)
	if err != nil {
		return err
	}

	return restErr
}
