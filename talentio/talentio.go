package talentio

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	libraryVersion = "1.0"
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
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.config.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return resp, err
}
