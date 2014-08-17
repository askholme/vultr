package vultr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const URL string = "https://api.vultr.com/v1"
const 
type Client struct {
	APIKey string
	URL string
  Http *htt.Client
  Params Parameters
}

var returncodes = map[int]string {
  200 = "OK"
  400 = "Invalid location (URL)"
  403 = "Invalid or missing API key"
  405 = "Invalid HTTP Method"
  500 = "Internal server error"
  412 = "ERROR"
}

func New(token string) (*Client, error) {
	// If it exists, grab teh token from the environment
	if apikey == "" {
		apikey = os.Getenv("VULTR_KEY")
	}

	client := Client{
		APIKey: apikey,
		Http:  http.DefaultClient,
	} 
  client.Params = initParameters(client)
	return &client, nil
}

// Creates a new request with the params
func (c *Client) Request(params map[string]string, action string,method string) (*http.Request, error) {
	p := url.Values{}
	u, err := url.Parse(c.URL + action)

	if err != nil {
		return nil, fmt.Errorf("Error parsing base URL: %s", err)
	}

	// Build up our request parameters
	for k, v := range params {
		p.Add(k, v)
	}
  p.add("api_key",c.APIKey)

	// Add the params to our URL
	u.RawQuery = p.Encode()

	// Build the request
	req, err := http.NewRequest(method, u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("Error creating request: %s", err)
	}

	return req, nil
}

// checkResp wraps http.Client.Do() and verifies that the
// request was successful. A non-200 request returns an error
// as per the vultr api documentation
func checkResp(resp *http.Response, err error) (*http.Response, error) {
	// If the err is already there, there was an error higher
	// up the chain, so just return that
	if err != nil {
		return resp, err
	}
	switch i := returncodes[resp.StatusCode]; {
	case i == "OK":
		return resp, nil
	case i == "ERROR":
		return nil, getBody(resp)
	default:
		return nil, fmt.Errorf("API Error: %s", i)
	}
}

func getBody(resp *http.Response, out interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return body
}