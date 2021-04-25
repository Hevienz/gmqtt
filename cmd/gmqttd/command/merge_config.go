package command

import (
	"github.com/DrmagicE/gmqtt/config"
	"github.com/DrmagicE/gmqtt/plugin/federation"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

func NewMergeConfigCommand() *cobra.Command {
	var outputPath string

	cmd := &cobra.Command{
		Use:   "mergeconfig",
		Short: "Merge config",
		Run: func(cmd *cobra.Command, args []string) {
			servers := os.Getenv("GMQTT_FEDERATION_SERVERS")

			if servers == "" {
				log.Fatal("env GMQTT_FEDERATION_SERVERS not set")
			}

			c, err := config.ParseConfig(ConfigFile)
			if os.IsNotExist(err) {
				must(err)
			} else {
				must(err)
			}

			fed, ok := c.Plugins["federation"]
			if !ok {
				log.Fatal("federation not set in plugins")
			}

			fed2, ok :=  fed.(*federation.Config)
			if !ok {
				log.Fatal("bad federation config")
			}

			fed2.RetryJoin = strings.Split(servers, " ")

			c.Plugins["federation"] = fed2

			d, err := yaml.Marshal(&c)
			if err != nil {
				log.Fatalf("yaml encode failed, %s", err)
			}

			err = os.WriteFile(outputPath, d, 0777)
			if err != nil {
				log.Fatalf("write file failed, %s", err)
			}
		},
	}
	cmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "gmqttd.yml", "The output file path")

	return cmd
}
