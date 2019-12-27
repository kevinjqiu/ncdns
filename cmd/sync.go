package cmd

import (
	"errors"
	"github.com/kevinjqiu/ncdns/pkg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func newSyncCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "sync",
		Short: "synchronize a zone",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("Must provide a DNS zone file")
			}

			zoneFileContent, err := ioutil.ReadFile(args[0])
			if err != nil {
				logrus.Fatal(err)
			}

			var syncConfig pkg.SyncConfig
			if err := yaml.Unmarshal(zoneFileContent, &syncConfig); err != nil {
				logrus.Fatal(err)
			}

			var ncDNS = pkg.NamecheapDNSUtil{}
			return nil
		},
	}

	return cmd
}
