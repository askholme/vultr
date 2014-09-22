package vultr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
  "encoding/json"
  
)
import . "github.com/visionmedia/go-debug"
var debug = Debug("Api")

type Client struct {
	APIKey string
	URL string
  Http *http.Client
  Params Parameters
}

var returncodes = map[int]string {
  200 : "OK",
  400 : "Invalid location (URL)",
  403 : "Invalid or missing API key",
  405 : "Invalid HTTP Method",
  500 : "Internal server error",
  412 : "ERROR",
}
func MakeClient(apikey string) (*Client, error) {
	// If it exists, grab teh token from the environment
	if apikey == "" {
		return nil,fmt.Errorf("no API key provided to vultr API")
	}
	client := Client{
		APIKey: apikey,
		Http:  http.DefaultClient,
    URL: "https://api.vultr.com/v1",
	}
  return &client,nil
}
func NewClient(apikey string) (*Client, error) {
  client,err := MakeClient(apikey)
  if err != nil {
    return nil,err
  }
  err = client.initParams()
  if err != nil {
    return nil,err
  }
  return client,nil
}
func (c *Client) initParams() (error) {
  params,err := NewParameters(c)
  if err != nil {
    return err
  }
  c.Params = params
  return nil
}
// send a request including api key
func (c *Client) RequestByte(params map[string]string, action string,method string) ([]byte, error) {
	p := url.Values{}
  p.Add("api_key",c.APIKey)
  // even on post we need api_key in the url
  url := ""
  if method == "POST" {
    url = c.URL + action + "?" + p.Encode()
  }
	// Build up our request parameters
	for k, v := range params {
		p.Add(k, v)
	}
	// Add the params to our URL
	if method == "GET" {
	  url = c.URL + action + "?" + p.Encode()
	}
  var body []byte
  var err error
  debug("will hit url %s\n",url)
  if method == "GET" {
    body, err = c.doReq(c.Http.Get(url))
  } else {
    body, err = c.doReq(c.Http.PostForm(url,p))
  }
	if err != nil {
		return nil,err
	}
	return body, nil
}
// create a new request and returns body as string
func (c *Client) RequestStr(params map[string]string, action string,method string) (string, error) {
  data,err := c.RequestByte(params,action,method)
  return string(data),err
}

func (c *Client) RequestInterface(params map[string]string, action string,method string,out interface{}) (error) {
  data,err := c.RequestByte(params,action,method)
  if err != nil {
    return err
  }
  err = json.Unmarshal(data, out)
  if err != nil {
    panic(fmt.Sprintf("Error reading json: %s", err))
  }
  return nil
}
// Create a new request and decodes to Jaons
func (c *Client) RequestMap(params map[string]string, action string,method string) (map[string]interface{}, error) {
  var f interface{} 
  err := c.RequestInterface(params,action,method,&f)
  if err != nil {
    return nil,err
  }
  switch test := f.(type) {
  case []interface{}:
    // got array which means empty
    return nil,nil
  case map[string]interface{}:
    return test,nil
  default:
    // unknown type
    panic(fmt.Sprintf("error parsing json"))
  }
}

// Create a new request and decodes to Jaons
func (c *Client) RequestArr(params map[string]string, action string,method string) ([]interface{}, error) {
  var f interface{} 
  err := c.RequestInterface(params,action,method,&f)
  if err != nil {
    return nil,err
  }
  switch test := f.(type) {
  case []interface{}:
    // got array 
    return test,nil
  case map[string]interface{}:
    panic(fmt.Sprintf("Got map while expecting array in JSON response"))
  default:
    // unknown type
    panic(fmt.Sprintf("error parsing json"))
  }
}


// checkResp wraps http.Client.Do() and verifies that the
// request was successful. A non-200 request returns an error
// as per the vultr api documentation
func (c *Client) doReq(resp *http.Response,err error) ([]byte,error) {
	// If the err is already there, there was an error higher
	// up the chain, so just return that
  //resp,err := c.Http.Do(req)
	if err != nil {
		return nil,fmt.Errorf("Error making request: %s",err)
	}
	switch i := returncodes[resp.StatusCode]; {
	  case i == "OK":
      body, err := c.getBody(resp)
      if err != nil {
        return nil,fmt.Errorf("Error reading body %s:", err)
      }
		  return body,nil
	  case i == "ERROR":
      body,err := c.getBody(resp)
      if err != nil {
        return nil,fmt.Errorf("API Error which could not be read")
      }
		  return nil,fmt.Errorf(string(body))
	  default:
		  return nil, fmt.Errorf("API Error: %s", i)
	}
}

func (c *Client) getBody(resp *http.Response) ([]byte, error) {
  defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}

	return body,nil
}