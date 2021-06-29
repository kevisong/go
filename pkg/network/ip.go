package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"
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

// GetIPAddrs gets private ip addresses
func GetIPAddrs() ([]net.Addr, error) {
	return net.InterfaceAddrs()
}

// GetIPAddrsAsString gets private ip addresses as string
func GetIPAddrsAsString() ([]string, error) {

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		errMsg := fmt.Sprintf("net.InterfaceAddrs() failed, error: %s", err)
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	ips := []string{}

	for _, address := range interfaceAddr {
		ipNet, isValidIPNet := address.(*net.IPNet)
		if isValidIPNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}

	return ips, nil

}
