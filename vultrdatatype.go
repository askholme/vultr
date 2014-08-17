package vultr

import {
	"fmt"
	"strconv"
	"strings"
}

type parameter struct {
 by_id  map[int]string
 by_label = map[string]int
}

type Parameters struct {
  Regions parameter `url:"regions"`
  Snapshosts parameter `url:"snapshots"`
  Images parameter `url:"images"`
  Plans parameter `url:"plans"`
  Sizes parameter `url:"sizes"`
}

func init() Parameters {
  p := Parameters{}
  
  return p
}