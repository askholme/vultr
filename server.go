package vultr

import ("time"
"strconv"
"fmt"
)
type Server struct {
  Id            int     `json:"SUBID"`
  Os            string  `json:"os"`
  Ram           string  `json:"ram"`
  Disk          string  `json:"disk"`
  Ip            string  `json:"main_ip"`
  Cpu           string  `json:"vcpu_count"`
  Location      string  `json:"location"`
  DCID          string  `json:"dcid"`
  Password      string  `json:"default_password"`
  Created       string  `json:"date_created"`
  Status        string  `json:"status"`
  Netmaskv4     string  `json:"netmask_v4"`
  Gatewayv4     string  `json:"gateway_v4"`
  PlanID        int     `json:"VPSPLANID"`
  IpV6          string  `json:"v6_main_ip"`
  NetmaskV6     string  `json:"v6_network"`
  SizeV6        string  `json:"v6_network_size"`
  PrivateIP     string  `json:"internal_ip"`
  Label         string  `json:"label"`
//  Charges       string  `json:"pending_charges"`
//  Cost          float64 `json:"cost_per_month"`
//  CurrentBW     float64 `json:"current_bandwidth_gb"`
//  AllowedVW     float64 `json:"alllowed_bandwidth_gb"`
  Power         string  `json:"power_status"`
  KVMurl        string  `json:"kvm_url"`            
}
type CreateServer struct {
  Region      string
  Plan        string
  Os          string
  Snapshot    string
  IpV6        bool
  PrivateNet  bool
  Name        string
  IpxeUrl     string
}
func (c *Client) CreateOpts() CreateServer {
  opts := CreateServer{}
  opts.Ipv6 = true
  opts.PrivateNet = true
  return opts
}
func (c *Client) TestRegionPlan(region_id string,plan_id string) (bool) {
  params := make(map[string]string)
  params["DCID"] = region_id
  //planid, _ := strconv.Atoi(plan_id)
  arr,err := c.RequestArr(params,"/regions/availability","GET")
  if err != nil {
    panic("error while checking region availability")
  }
  // run throug array and return if plan found
  for _, v := range arr {
    switch test := v.(type) {
    case string:
      if test == plan_id { return true }
    case int:
      if strconv.Itoa(test) == plan_id { return true }
    case float64:
      if strconv.FormatFloat(test,'f',-1,64) == plan_id { return true }
    default:
      panic(fmt.Sprintf("error parsing region availabilty %s",test))
    }
  }
  return false
}
func (c *Client) CreateServer(opts *CreateServer) (string,error) {
  params := make(map[string]string)
  // the get id function handles conversion to Id if needed
  var err error
  params["DCID"],err = c.Params.getId("region",opts.Region)
  if err != nil {
    return "",err
  }
  params["VPSPLANID"],err = c.Params.getId("plan",opts.Plan)
  if err != nil {
    return "",err
  }
  if !c.TestRegionPlan(params["DCID"],params["VPSPLANID"]) {
    return "",fmt.Errorf("Plan %s not available in region %s",opts.Plan,opts.Region)
  }
  osOptions := 0
  if (opts.Os != "") { osOptions++ }
  if (opts.Snapshot != "") { osOptions++ }
  if (opts.IpXeUrl != "") { osOptions++ }
  if osOptions>1 {
    return "",fmt.Errorf("OS, Snapshot and Ixpe parameters cannot be combined")
  }
  if opts.Snapshot != "" {
    params["SNAPSHOTID"],err = c.Params.getId("snapshot",opts.Snapshot)
    if err != nil {
      return "",err
    }
    params["OSID"],err = c.Params.getId("os","Snapshot")
  } else if opts.IpxeUrl != "" {
    // tests for this and snapshot handling would be good
    params["ipxe_chain_url"] = opts.IpxeUrl
    params["OSID"],err = c.Params.getId("os","Custom")
  } else {
    params["OSID"],err = c.Params.getId("os",opts.Os)
  }
  if err != nil {
    return "",err
  }
  if opts.IpV6 {
    params["enable_ipv6"] = "yes"
  }
  if opts.PrivateNet {
    params["enable_private_network"] = "yes"
  }
  if opts.Name != "" {
    params["label"] = opts.Name
  }
  resp,err := c.RequestMap(params,"/server/create","POST")
  if err != nil {
    return "",fmt.Errorf("Error executing request %s",err)
  }
  val := resp["SUBID"].(string)
  return val,nil
}
func (c *Client) GetServers() (map[string]Server,error) {
  serverlist := make(map[string]Server)
  err := c.RequestInterface(nil,"/server/list","GET",&serverlist)
  if err != nil {
    return serverlist,err
  }
  return serverlist,nil
}
func (c *Client) GetServer(search_id string) (*Server,error) {
  serverlist,err := c.GetServers()
  if err != nil {
    return nil,err
  }
  //var s Server
  for id,server := range serverlist {
    if id == search_id {
      debug("%s did match in servers list",id)
      return &server,nil
    }
    debug("%s did not match in servers list",id)
  }
  return nil,fmt.Errorf("Server not found")
}
func (c *Client) WaitForServer(id string) (*Server,error) {
  for i := 0; i < 150; i++ {
    server,err := c.GetServer(id)
    if err != nil {
      return nil,err
    }
    if server != nil {
      return server,nil
    }
    time.Sleep(2 * time.Second)
  }
  panic("timeout while waiting for server to boot")
}
func (c *Client) DeleteServer(id string) (error) {
  params := make(map[string]string)
  params["SUBID"] = id
  _, err := c.RequestStr(params,"/server/destroy","POST")
  return err
}
func (c *Client) SetServerIpV4Reverse(id string,ip string, dns string) (error) {
  params := make(map[string]string)
  params["SUBID"] = id
  params["ip"] = ip
  params["entry"] = dns
  _, err := c.RequestStr(params,"/server/reverse_set_ipv4","POST")
  return err
}
func (c *Client) SetServerIpV6Reverse(id string,ip string, dns string) (error) {
  params := make(map[string]string)
  params["SUBID"] = id
  params["ip"] = ip
  params["entry"] = dns
  _, err := c.RequestStr(params,"/server/reverse_set_ipv6","POST")
  return err
}
func (c *Client) SetServerLabel(id string,label string) (error) {
  params := make(map[string]string)
  params["SUBID"] = id
  params["label"] = label
  _, err := c.RequestStr(params,"/server/label_set","POST")
  return err
}
