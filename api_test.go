package vultr 
import (
	"github.com/pearkes/digitalocean/testutil"
  "testing"
)
import 	. "github.com/motain/gocheck"
type S struct {
	client *Client
}
func makeResp(body string) testutil.Response {
  resp := testutil.Response{}
  resp.Status = 200
  resp.Headers = make(map[string]string)
  resp.Headers["Content-type"] = "application/json"
  resp.Body = body
  return resp
}
var _ = Suite(&S{})
var testServer = testutil.NewHTTPServer()
var paramResponses = testutil.ResponseMap {
  "/v1/os/list" : makeResp(`{
        "127": {
            "OSID": "127",
            "name": "CentOS 6 x64",
            "arch": "x64",
            "family": "centos",
            "windows": false
        },
        "148": {
            "OSID": "148",
            "name": "Ubuntu 12.04 i386",
            "arch": "i386",
            "family": "ubuntu",
            "windows": false
        }
    }`),
  "/v1/plans/list": makeResp(`{
      "1": {
          "VPSPLANID": "1",
          "name": "Starter",
          "vcpu_count": "1",
          "ram": "512",
          "disk": "20",
          "bandwidth": "1",
          "price_per_month": "5.00",
          "windows": false
      },
      "2": {
          "VPSPLANID": "2",
          "name": "Basic",
          "vcpu_count": "1",
          "ram": "1024",
          "disk": "30",
          "bandwidth": "2",
          "price_per_month": "8.00",
          "windows": false
      }
  }`),
  "/v1/regions/list": makeResp(`{
        "1": {
            "DCID": "1",
            "name": "New Jersey",
            "country": "US",
            "continent": "North America",
            "state": "NJ"
        },
        "2": {
            "DCID": "2",
            "name": "Chicago",
            "country": "US",
            "continent": "North America",
            "state": "IL"
        }
    }`),
  "/v1/snapshot/list": makeResp(`{
      "5359435d28b9a": {
          "SNAPSHOTID": "5359435d28b9a",
          "date_created": "2014-04-18 12:40:40",
          "description": "Test snapshot",
          "size": "42949672960",
          "status": "complete"
      },
      "5359435dc1df3": {
          "SNAPSHOTID": "5359435dc1df3",
          "date_created": "2014-04-22 16:11:46",
          "description": "",
          "size": "10000000",
          "status": "complete"
      }
  }`),
}
func TestServer(t *testing.T) {
	TestingT(t)
}
func (s *S) SetUpSuite(c *C) {
	testServer.Start()
	var err error
	s.client, err = MakeClient("foobar")
	if err != nil {
		panic(err)
	}
  s.client.URL = "http://localhost:4444/v1"
  // configure test server to send parameters
  testServer.ResponseMap(4,paramResponses)
  err = s.client.initParams()
  testServer.WaitRequests(4)
  // pull the 4 requests out so they don't mess up other stuff
	if err != nil {
		panic(err)
	}
}
func (s *S) TearDownTest(c *C) {
	testServer.Flush()
}