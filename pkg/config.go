package pkg

import (
	"errors"
	"fmt"
	"strings"
)

type (
	ApiConfig struct {
		APIUser  string `yaml:"apiUser"`
		Username string `yaml:"username"`
		Token    string `yaml:"token"`
	}

	Config struct {
		API ApiConfig `yaml:"api"`
	}

	SyncRecordConfig struct {
		Type      string `yaml:"type"`
		Name      string `yaml:"name"`
		Address   string `yaml:"address"`
		TTL       int    `yaml:"ttl"`
		CreatePTR bool   `yaml:"createPTR"`
	}

	SyncConfig struct {
		Zone    string             `yaml:"zone"`
		Records []SyncRecordConfig `yaml:"records"`
	}
)

func (src SyncRecordConfig) FQDN(zone string) string {
	return src.Name + "." + zone
}

func (src SyncRecordConfig) PTR() (string, error) {
	if src.Type != "A" {
		return "", errors.New("record type must be A in order to create PTR records")
	}

	octets := strings.Split(src.Address, ".")
	if len(octets) != 4 {
		return "", fmt.Errorf("invalid address of type A: %s", src.Address)
	}

	return fmt.Sprintf("%s.%s.%s.%s.in-addr.arpa", octets[3], octets[2], octets[1], octets[0]), nil
}