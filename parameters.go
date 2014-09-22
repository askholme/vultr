package vultr

import (
	"fmt"
)
type Parameter struct {
 by_id  map[string]string
 by_label map[string]string
}

type Parameters struct {
  Params      map[string]Parameter
  client      *Client
}
func (p *Parameters) getParam(paramname string) (Parameter,error) {
  param, ok := p.Params[paramname]
  if !ok {
    return Parameter{},fmt.Errorf("Parameter %s is not initialized",paramname)
  }
  return param,nil
}
func (p *Parameters) getId(paramname string,label string) (string,error) {
  param, err := p.getParam(paramname)
  if err != nil {
    return "",err
  }
  // if label is and id which is known then just throw it bac
  if isInt(label) {
    _, ok := param.by_id[label]
    if ok {
      return label,nil
    }
  }
  val,ok  := param.by_label[label]
  if !ok {
    return "",fmt.Errorf("Label %s not found in parameter %s",label,paramname)
  }
  return val,nil
}
func (p *Parameters) getLabel(paramname string,id string) (string,error) {
  param, err := p.getParam(paramname)
  if err != nil {
    return "",err
  }
  val,ok  := param.by_id[id]
  if !ok {
    return "",fmt.Errorf("Id %s not found in parameter %s",id,paramname)
  }
  return val,nil
}
func (p *Parameters) initParam(name string, url string, namefield string) (error) {
  param := Parameter{}
  param.by_id = make(map[string]string)
  param.by_label = make(map[string]string)
  data,err := p.client.RequestMap(nil, fmt.Sprintf("/%s/list", url), "GET")
  if err != nil {
    return err
  }
  for id,arr := range data {
    real_map := arr.(map[string]interface{})
    for key,value := range real_map {
      if key == namefield {
        val := value.(string)
        param.by_id[id] = val
        if val != "" {
            //snapshots might not have a label
            param.by_label[value.(string)] = id
        }
        debug("%s - %s:%s:%s\n",name,key,value,id)
      }  
    }
  }
  debug("iniialized param %s",name)
  p.Params[name] = param
  return nil
}
func NewParameters(client *Client) (Parameters,error) {
  p := Parameters{}
  p.client = client
  p.Params = make(map[string]Parameter)
  var err error
  err = p.initParam("region","regions","name")
  if err != nil { return p,err }
  err = p.initParam("plan","plans","name")
  if err != nil { return p,err }
  err = p.initParam("os","os","name")
  if err != nil { return p,err }
  err = p.initParam("snapshot","snapshot","description")   
  if err != nil { return p,err }
  return p,nil
}