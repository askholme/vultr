package vultr

import {
	"fmt"
	"strconv"
	"strings"
  "reflect"
}

type parameter struct {
 by_id  map[int]string
 by_label map[string]int
}

type Parameters struct {
  Regions     parameter `url:"regions        id:"DCID"        name:"name"`
  Snapshosts  parameter `url:"snapshots"     id:"SNAPSHOTID"  name:"description"`
  Images      parameter `url:"images"        id:"IMAGEID"     name:"name"`
  Plans       parameter `url:"plans"         id:"VPSPLANID"   name:"name"`
  Scripts     parameter `url:"startupscript" id:"SCRIPTID"    name:"name"`
  OS          parameter `url:"os"            id:"OSID"        name:"name"`
}

func initParameters(Client client) Parameters {
  p := Parameters{}
  p_type_val := reflect.ValueOf(&p).Elem()
  p_type := reflect.TypeOf(p)
  for i:= 0; i< p_type.NumField(); i++ {
    f := p_type_val.Field(i)
    name = p_type.Field(i).Name
    action = p_type.Field(i).Tag.get("url")
    fmt.Printf("%d: %s %s",i,name,tag)
    data = client.Request(nil, "GET", fmt.SprintF("/%s/list", url))
    // Parse data to get id and name colmn. Note that it's a hash of hashes, so first key is id and then ID opens for something more
  }
  return p
}