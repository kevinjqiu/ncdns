package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/billputer/go-namecheap"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ncdns",
	Short: "Namecheap DNS CLI utility",
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	//Run: func(cmd *cobra.Command, args []string) {
	//	apkToken := "xxxx"
	//	c := namecheap.NewClient("kevinjqiu", apkToken, "kevinjqiu")
	//	//domains, _ := c.DomainsGetList()
	//	//for _, domain := range domains {
	//	//	fmt.Printf("Domain: %+v\n\n", domain.Name)
	//	//}
	//	//
	//	hostRecords, err := c.DomainsDNSGetHosts("idempotent", "io")
	//	hosts := hostRecords.Hosts
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	r := findRecord(hosts, "satellite")
	//	if r == nil {
	//		r = &namecheap.DomainDNSHost{
	//			Name:    "satellite",
	//			Type:    "A",
	//			Address: "192.168.200.46",
	//			TTL:     300,
	//		}
	//	}
	//	hosts = append(hosts, *r)
	//	result, err := c.DomainDNSSetHosts("idempotent", "io", hosts)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(result)
	//},
}

func findRecord(records []namecheap.DomainDNSHost, hostname string) *namecheap.DomainDNSHost {
	for _, r := range records {
		if r.Name == hostname {
			return &r
		}
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ncdns.yaml)")
	rootCmd.AddCommand(newSyncCommand())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ncdns" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ncdns")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}