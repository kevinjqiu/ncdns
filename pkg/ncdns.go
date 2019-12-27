package pkg

import "github.com/billputer/go-namecheap"

type NamecheapDNSUtil struct {
	client *namecheap.Client
}

func (ncdns NamecheapDNSUtil) Sync(config SyncConfig) error {
	return nil
}