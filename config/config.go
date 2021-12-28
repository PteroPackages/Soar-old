package config

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var GroupCommand = &cobra.Command{
	Use: "config",
}

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Gets information about the Soar configuration.",
	Long:  "Gets information about the Soar configuration.",
	Run: func(cmd *cobra.Command, _ []string) {
		config, err := GetConfig()
		if err != nil {
			fmt.Printf("soar: %s", err.Error())
			return
		}

		var out strings.Builder

		out.WriteString("Soar Config\n")
		out.Write([]byte(strings.Repeat("=", 20)))
		out.WriteString("\napplication API\n")
		out.WriteString(fmt.Sprintf("\turl: %s\n", config.Application.URL))
		out.WriteString(fmt.Sprintf("\tkey: %s\n", config.Application.Key))
		out.WriteString("\n\nclient API\n")
		out.WriteString(fmt.Sprintf("\turl: %s\n", config.Client.URL))
		out.WriteString(fmt.Sprintf("\tkey: %s\n", config.Client.Key))

		fmt.Println(out.String())
	},
}

func init() {
	GroupCommand.AddCommand(InfoCmd)
}
