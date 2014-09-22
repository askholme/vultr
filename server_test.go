package vultr 
import (
	"github.com/pearkes/digitalocean/testutil"
)

import . "github.com/motain/gocheck"
var createResponses = testutil.ResponseMap{
  "/v1/regions/availability": makeResp("[1]"),
  "/v1/server/create": makeResp(`{ "SUBID" : "576965" }`),
}
var getResponses = testutil.ResponseMap{
  "/v1/server/list": makeResp(`{
      "576965": {
          "SUBID": 576965,
          "os": "CentOS 6 x64",
          "ram": "4096 MB",
          "disk": "Virtual 60 GB",
          "main_ip": "123.123.123.123",
          "vcpu_count": "2",
          "location": "New Jersey",
          "DCID": "1",
          "default_password": "nreqnusibni",
          "date_created": "2013-12-19 14:45:41",
          "pending_charges": "46.67",
          "status": "active",
          "cost_per_month": 10.05,
          "current_bandwidth_gb": 131.512,
          "allowed_bandwidth_gb": 1000,
          "netmask_v4": "255.255.255.248",
          "gateway_v4": "123.123.123.1",
          "power_status": "running",
          "VPSPLANID": 28,
          "v6_network": "2001:DB8:1000::",
          "v6_main_ip": "2001:DB8:1000::100",
          "v6_network_size": "64",
          "label": "my new server",
          "internal_ip": "10.99.0.10",
          "kvm_url": "https://my.vultr.com/subs/novnc/api.php?data=eawxFVZw2mXnhGUV"
      }
    }`)}

func (s *S) Test_CreateServer_1(c *C) {
  testServer.ResponseMap(2,createResponses)
  opts := s.client.CreateOpts()
  opts.Region = "New Jersey"
  opts.Plan = "Starter"
  opts.Os = "Ubuntu 12.04 i386"
  id,err := s.client.CreateServer(&opts)
  reqs := testServer.WaitRequests(2)
  c.Assert(err, IsNil)
  c.Assert(id,Equals,"576965")
  c.Assert(reqs[1].Form.Get("VPSPLANID"),Equals,"1")
  c.Assert(reqs[1].Form.Get("OSID"),Equals,"148")
  c.Assert(reqs[0].Form.Get("DCID"),Equals,"1")
}
func (s *S) Test_CreateServer_2(c *C) {
  testServer.ResponseMap(2,createResponses)
  opts := s.client.CreateOpts()
  opts.Region = "New Jersey"
  opts.Plan = "Basic"
  opts.Os = "Ubuntu 12.04 i386"
  _,err := s.client.CreateServer(&opts)
  c.Assert(err, ErrorMatches, ".*not available in region.*")
}

func (s *S) Test_GetServer(c *C) {
  testServer.ResponseMap(1,getResponses)
  server,err := s.client.GetServer("576965")
  if err != nil{
    panic(err)
  }
  c.Assert(err,IsNil)
  c.Assert(server,Not(IsNil))
  c.Assert(server.Ram,Equals,"4096 MB")
  c.Assert(server.PrivateIP,Equals,"10.99.0.10")
}
