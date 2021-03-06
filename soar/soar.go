package soar

import (
	"log"

	"github.com/pteropackages/soar/app"
	"github.com/pteropackages/soar/cmd"
	"github.com/pteropackages/soar/config"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "soar",
	Short: "Connect to Pterodactyl instances",
	Long:  "Connect to a Pterodactyl instance locally or remotely via command line",
}

func init() {
	root.AddCommand(cmd.VersionCmd)
	root.AddCommand(cmd.InfoCmd)
	root.AddCommand(config.GroupCommand)
	root.AddCommand(app.GroupCommand)
}

func Execute() {
	if err := root.Execute(); err != nil {
		log.Fatalf("Failed running Soar: %s", err)
	}
}
