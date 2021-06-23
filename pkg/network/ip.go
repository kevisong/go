package network

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
)

// IPInfo defines ip information
type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	LOC      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

// GetPublicIPInfo gets public ip information
func GetPublicIPInfo() (IPInfo, error) {

	ipInfo := IPInfo{}

	resp, err := http.Get("http://ipinfo.io/")
	if err != nil {
		return ipInfo, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ipInfo, err
	}

	err = json.Unmarshal(respBytes, &ipInfo)
	if err != nil {
		return ipInfo, err
	}
	return ipInfo, nil

}

// GetIPAddrs gets private ip information from net interface
func GetIPAddrs() ([]net.Addr, error) {
	return net.InterfaceAddrs()
}
