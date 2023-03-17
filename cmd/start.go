package cmd

import (
	"io/ioutil"

	"github.com/golang-code-specification/internal/app"
	"github.com/golang-code-specification/internal/app/config"
	"github.com/spf13/cobra"
)

const (
	defaultLocalConfig = "/opt/local.yaml"
)

var (
	localConfig string
	startCmd    = &cobra.Command{
		Use:   "start",
		Short: "Start Project Server.",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&localConfig, "CONFIG", "c", defaultLocalConfig, "conf path: /opt/local.yaml")
}

func start() {
	conf := localConf()
	app.Serve(conf)
}

func localConf() *config.Config {
	if localConfig == "" {
		localConfig = defaultLocalConfig
	}
	data, err := ioutil.ReadFile(localConfig)
	if err != nil {
		panic(err)
	}
	conf, err := config.ReadConfig(data)
	if err != nil {
		panic(err)
	}
	return conf
}
