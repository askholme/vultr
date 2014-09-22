package vultr
import (
  "fmt"
)
type Snapshot struct {
  Id            string  `json:"SNAPSHOTID"`
  Created       string  `json:"date_created"`
  Description   string  `json:"description"`
  Size          string  `json:"size"`
  Status        string  `json:"status"`
}

func (c *Client) CreateSnapshot(serverId string,description string) (string,error) {
  params := make(map[string]string)
  params["SUBID"] = serverId
  params["description"] = description
  data := make(map[string]string)
  err := c.RequestInterface(params,"/snapshot/create","POST",&data)
  if err != nil {
    return "",err
  }
  return data["SNAPSHOTID"],err
}
func (c *Client) DeleteSnapshot(snapshotId string) error {
  params := make(map[string]string)
  params["SUBID"] = snapshotId
  _, err := c.RequestStr(params,"/snapshot/destroy","POST")
  return err
}
func (c *Client) GetSnapshots() (map[string]Snapshot,error) {
  snapshotlist := make(map[string]Snapshot)
  err := c.RequestInterface(nil,"/snapshot/list","GET",&snapshotlist)
  if err != nil {
    return snapshotlist,err
  }
  return snapshotlist,nil
}
func (c *Client) GetSnapshot(snapshotId string) (*Snapshot,error) {
  serverlist,err := c.GetSnapshots()
  if err != nil {
    return nil,err
  }
  //var s Server
  for id,snapshot := range serverlist {
    if id == snapshotId {
      debug("%s did match in snapshot list",id)
      return &snapshot,nil
    }
    debug("%s did not match in snapshot list",id)
  }
  return nil,fmt.Errorf("Snapshot not found")
}