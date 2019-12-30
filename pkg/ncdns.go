package pkg

import (
	"github.com/billputer/go-namecheap"
	"github.com/sirupsen/logrus"
	"strings"
)

type NamecheapDNSUtil struct {
	client *namecheap.Client
}

type zone string

func (z zone) split() [2]string {
	parts := strings.Split(string(z), ".")
	sld := strings.Join(parts[0:len(parts)-1], ".")
	tld := parts[len(parts)-1]
	return [2]string{sld, tld}
}

func (z zone) SLD() string {
	return z.split()[0]
}

func (z zone) TLD() string {
	return z.split()[1]
}

func (ncdns NamecheapDNSUtil) findExistingRecord(existingHosts []namecheap.DomainDNSHost, r SyncRecordConfig) (*namecheap.DomainDNSHost, bool) {
	for _, existingHost := range existingHosts {
		logrus.Info(existingHost)
		if existingHost.Name == r.Name && existingHost.Type == r.Type {
			return &existingHost, true
		}
	}
	return nil, false
}

func (ncdns NamecheapDNSUtil) Sync(config SyncConfig) error {
	zone := zone(config.Zone)

	var (
		sld = zone.SLD()
		tld = zone.TLD()
	)
	logrus.Infof("sld=%s, tld=%s", sld, tld)
	getResult, err := ncdns.client.DomainsDNSGetHosts(sld, tld)
	if err != nil {
		return err
	}

	hosts := getResult.Hosts
	var toAdd []namecheap.DomainDNSHost
	for _, desiredRecord := range config.Records {
		existingHost, ok := ncdns.findExistingRecord(hosts, desiredRecord)
		if !ok {
			toAdd = append(toAdd, namecheap.DomainDNSHost{
				Name:    desiredRecord.Name,
				Type:    desiredRecord.Type,
				Address: desiredRecord.Address,
				TTL:     desiredRecord.TTL,
			})
		} else {
			existingHost.Address = desiredRecord.Address
			existingHost.TTL = desiredRecord.TTL
		}

		if desiredRecord.CreatePTR {
			ptr, err := desiredRecord.PTR()
			if err != nil {
				logrus.Warn("Unable to create PTR record: ", err)
			}

			ptrRecord, ok := ncdns.findExistingRecord(hosts, SyncRecordConfig{
				Type: "PTR",
				Name: ptr,
			})

			if !ok {
				toAdd = append(toAdd, namecheap.DomainDNSHost{
					Name:    ptr,
					Type:    "PTR",
					Address: desiredRecord.FQDN(string(zone)),
					TTL:     desiredRecord.TTL,
				})
			} else {
				ptrRecord.Address = desiredRecord.FQDN(string(zone))
			}
		}
	}

	logrus.Info("New records to add:")
	for _, newHost := range toAdd {
		logrus.Info(newHost)
		hosts = append(hosts, newHost)
	}

	setResult, err := ncdns.client.DomainDNSSetHosts(sld, tld, hosts)
	if err != nil {
		logrus.Warn(setResult)
		return err
	}
	logrus.Info(setResult)
	logrus.Info("OK")
	return nil
}

func NewNamecheapDNSUtil(config Config) (*NamecheapDNSUtil, error) {
	return &NamecheapDNSUtil{
		client: namecheap.NewClient(config.API.APIUser, config.API.Token, config.API.Username),
	}, nil
}
