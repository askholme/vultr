package vultr

type Server struct {
  id            int `json:"SUBID"`
  os            string `json:"os"`
  ram           string `json:"ram"`
  disk          string `json:"disk"`
  ip            string `json:"main_ip"`
  cpu           string `json:"vcpu_count"`
  location      string `json:"location"`
  dcid          string `json:"dcid"`
  password      string `json:"default_password"`
  created       string `json:"created_date"`
  status        string `json:"status"`
  netmaskv4     string `json:"netmask_v4"`
  gatewayv4     string `json:"gateway_v4"`
  pricing_plan  int `json:"VPSPLANID"`
  ipv6          string `json:"v6_main_ip"`
  netmaskv6     string `json:"v6_network"`
  private_ip    string `json:"internal_ip"`
  label         string `json:"label"`           
}