package vultr


import . "github.com/motain/gocheck"
func (s *S) Test_Parameters(c *C) {
  str, err := s.client.Params.GetLabel("region","1")
  c.Assert(err,IsNil)
  c.Assert(str,Equals,"New Jersey")
  str, err = s.client.Params.GetLabel("os","127")
  c.Assert(err,IsNil)
  c.Assert(str,Equals,"CentOS 6 x64")
  str,err = s.client.Params.GetLabel("plan","1")
  c.Assert(err,IsNil)
  c.Assert(str,Equals,"Starter")
  str,err = s.client.Params.GetId("snapshot","Test snapshot")
  c.Assert(err,IsNil)
  c.Assert(str,Equals,"5359435d28b9a")
  str,err = s.client.Params.GetId("snapshot","foo")
  c.Assert(err,ErrorMatches,".*not found in parameter.*")
}