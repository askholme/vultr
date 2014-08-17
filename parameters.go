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
  Regions parameter `url:"/regions"`
  Snapshosts parameter `url:"/snapshots"`
  Images parameter `url:"/images"`
  Plans parameter `url:"/plans"`
  Sizes parameter `url:"/sizes"`
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
    client.Request(nil, "GET", fmt.SprintF("/%s/list", url))
  }
  return p
}