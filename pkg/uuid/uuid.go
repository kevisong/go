package uuid

import (
	"errors"
	"fmt"
	"hash/crc32"

	"github.com/KEVISONG/go/pkg/network"
	log "github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

// GetIPListHash16 GetIPListHash16
func GetIPListHash16(ipList []string) uint16 {
	crc32q := crc32.MakeTable(0xD5828281)
	hash32 := crc32.Checksum([]byte(fmt.Sprintf("%s", ipList)), crc32q)
	return uint16(hash32)
}

// Init initialize generator
func Init() error {

	ipList, err := network.GetIPAddrsAsString()
	if err != nil {
		errMsg := fmt.Sprintf("network.GetIPAddrsAsString() failed, error: %s", err)
		log.Error(errMsg)
		return errors.New(errMsg)
	}

	st := sonyflake.Settings{}
	st.MachineID = func() (uint16, error) { return GetIPListHash16(ipList), nil }
	st.CheckMachineID = func(uint16) bool { return true }
	sf = sonyflake.NewSonyflake(st)

	return nil
}

// Generate generates an UUID
func Generate() (string, error) {
	id, err := sf.NextID()
	if err != nil {
		errorInfo := fmt.Sprintf("Generate ID failed, error: %s", err)
		log.Error(errorInfo)
		return "", errors.New(errorInfo)
	}
	return fmt.Sprintf("%x", id), nil
}
