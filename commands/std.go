package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Shows the current version of Soar",
	Long:    "Shows the current version of Soar as well as contributors",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Printf("Soar version %s (%s)\n\n© 2021 devnote-dev", "0.0.1a", "27/11/2021 02:39")
	},
}
