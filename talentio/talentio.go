package talentio

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strconv"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

const (
	libraryVersion = "1.1"
	apiVersion     = "v1"
	defaultBaseURL = "https://talentio.com/api/"
	userAgent      = "talentio/" + libraryVersion

	headerXRemaining = "X-REMAINING"
	headerXReset     = "X-RESET"
	headerXTotal     = "X-TOTAL"
)

// A Client manages communication with the Talentio API.
type Client struct {
	config    *Config
	baseURL   *url.URL
	userAgent string
	common    service

	Candidates *CandidatesService
}

type service struct {
	client *Client
}

// Response is a Talentio API response. This wraps the standard http.Response
// returned from Talentio and provides convenient access to things.
type Response struct {
	*http.Response

	Remaining int
	Reset     int
	Total     int
}

// NewClient returns a new Talentio client. If a nil Config is
// provided, defaultConfig will be used.
func NewClient(config *Config) *Client {
	if config == nil {
		c := defaultConfig
		config = &c
	}
	if config.httpClient == nil {
		config.httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		config:    config,
		baseURL:   baseURL,
		userAgent: userAgent,
	}

	c.common.client = c
	c.Candidates = (*CandidatesService)(&c.common)
	return c
}

// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewRequest returns a new API Request given a method, URL, and optional body.
// A relative URL can be provided in urlStr, in which case it is resolved
// relative to the BaseURL of the Client. Relative URLs should always be
// specified without a preceding slash.
func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(path.Join(apiVersion, urlStr))
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.config.accessToken)
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.config.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			err = errors.New(string(b))
		}
		return nil, err
	}

	response := newResponse(resp)
	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		switch err {
		case io.EOF:
			// ignore EOF errors caused by empty response body
			err = nil
		}
	}

	return response, err
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	resp := &Response{
		Response: r,
	}
	if v := r.Header.Get(headerXRemaining); v != "" {
		resp.Remaining, _ = strconv.Atoi(v)
	}
	if v := r.Header.Get(headerXReset); v != "" {
		resp.Reset, _ = strconv.Atoi(v)
	}
	if v := r.Header.Get(headerXTotal); v != "" {
		resp.Total, _ = strconv.Atoi(v)
	}
	return resp
}
