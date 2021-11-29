package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Shows the current version of Soar",
	Long:    "Shows the current version of Soar as well as contributors",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Printf("Soar version %s", "0.0.1a")
	},
}

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Shows detailed information about Soar",
	Long:  "Shows detailed information about Soar",
	Run: func(cmd *cobra.Command, _ []string) {
		people := []string{
			"Devonte <https://github.com/devnote-dev>",
		}
		fmt.Printf("Soar Info\n\ncopyright (c) 2021 PteroPackages\ncontributors:\n- %s", strings.Join(people, "\n- "))
	},
}
