package vultr

import (
	"github.com/pearkes/digitalocean/testutil"
)
import . "github.com/motain/gocheck"

var snapshot_createResponses = testutil.ResponseMap{
  "/v1/snapshot/create": makeResp(`{ "SNAPSHOTID": "5359435d28b9a" }`),
}
var snapshot_getResponses = testutil.ResponseMap{
  "/v1/snapshot/list": makeResp(`{ 
  "5421de1839f36" : { "SNAPSHOTID" : "5421de1839f36",
      "date_created" : "2014-09-23 16:54:48",
      "description" : "test",
      "size" : "16106127360",
      "status" : "complete"
    },
  "5422e5396566a" : { "SNAPSHOTID" : "5422e5396566a",
      "date_created" : "2014-09-24 11:37:29",
      "description" : "mysnapsnot-1411572827",
      "size" : "16106127360",
      "status" : "complete"
    }
}`)}

func (s *S) Test_CreateSnapshot(c *C) {
  testServer.ResponseMap(1,snapshot_createResponses)
  id,err := s.client.CreateSnapshot("576965","Test snapshot")
  reqs := testServer.WaitRequests(1)
  c.Assert(err, IsNil)
  c.Assert(id,Equals,"5359435d28b9a")
  c.Assert(reqs[0].Form.Get("SUBID"),Equals,"576965")
}

func (s *S) Test_GetSnapshot(c *C) {
  testServer.ResponseMap(1,snapshot_getResponses)
  snap,err := s.client.GetSnapshot("5422e5396566a")
  c.Assert(err, IsNil)
  c.Assert(snap.Size,Equals,"16106127360")
}
func (s *S) Test_GetSnapshot_err(c *C) {
  testServer.ResponseMap(1,snapshot_getResponses)
  _,err := s.client.GetSnapshot("5359435dc1dd3")
  c.Assert(err, ErrorMatches, "Snapshot not found")
}