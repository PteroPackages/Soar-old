package soar

import (
	"log"

	"github.com/pteropackages/soar/commands"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "soar",
	Short: "Connect to Pterodactyl instances",
	Long:  "Connect to a Pterodactyl instance locally or remotely via command line",
}

func init() {
	root.AddCommand(commands.VersionCmd)
	root.AddCommand(commands.ConfigCmd)
}

func Execute() {
	if err := root.Execute(); err != nil {
		log.Fatalf("Failed running Soar: %s", err)
	}
}
