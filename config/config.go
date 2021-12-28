package config

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var GroupCommand = &cobra.Command{
	Use: "config",
}

var infoCmd = &cobra.Command{
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

var setupCommand = &cobra.Command{
	Use:   "setup",
	Short: "Sets up the Soar configuration directories and files.",
	Long:  "Sets up the Soar configuration directories and files.",
	Run: func(cmd *cobra.Command, _ []string) {
		_, err := GetConfig()
		if err == nil {
			fmt.Print("Config file already found, do you want to reset? (y/n) ")
			if !loopInputA() {
				fmt.Println("Cancelling setup...")
				return
			}
		}

		fmt.Print("\nWhere should the files be setup?\nLeave blank to use the default path: ")
		fp := loopInputB()
		if fp == "" {
			fmt.Println("Using default Soar path instead")
			if runtime.GOOS == "windows" {
				fp = "C:\\soar"
			} else {
				fp = "/soar"
			}
		}

		err = CreateEnv(fp)
		if err != nil {
			fmt.Printf("soar: %s", err.Error())
			return
		}

		fmt.Printf("Successfully setup soar directories at: %s", fp)
	},
}

func loopInputA() bool {
	var res string
	for {
		fmt.Scanln(&res)
		switch strings.ToLower(res) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			continue
		}
	}
}

func loopInputB() string {
	var res string
	for {
		fmt.Scanln(&res)
		if res == "" {
			return ""
		}
		if Exists(res) {
			return res
		}
	}
}

func init() {
	GroupCommand.AddCommand(infoCmd)
	GroupCommand.AddCommand(setupCommand)
}
