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
  snap,err := s.client.GetSnapshot("5359435dc1df3")
  c.Assert(err, IsNil)
  c.Assert(snap.Size,Equals,"10000000")
}
func (s *S) Test_GetSnapshot_err(c *C) {
  testServer.ResponseMap(1,snapshot_getResponses)
  _,err := s.client.GetSnapshot("5359435dc1dd3")
  c.Assert(err, ErrorMatches, "Snapshot not found")
}