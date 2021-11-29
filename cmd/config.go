package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type LogConfig struct {
	StrictMode       bool   `yaml:"strict_mode"`
	ShowDebug        bool   `yaml:"show_debug"`
	ShowHTTPLog      bool   `yaml:"show_http_log"`
	ShowWebSocketLog bool   `yaml:"show_ws_log"`
	ErrorOutDir      string `yaml:"error_out_dir"`
}

type SoarConfig struct {
	Version string `yaml:"version"`

	Application struct {
		URL string `yaml:"url"`
		Key string `yaml:"key"`
	} `yaml:"application"`

	Client struct {
		URL string `yaml:"url"`
		Key string `yaml:"key"`
	} `yaml:"client"`

	Logs LogConfig `yaml:"logs"`
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Soar config",
	Long:  "Manages the internal configuration for Soar",
}

// File path checker
// Thanks stackoverflow!
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

var configSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the configuration file for Soar",
	Long:  "Setup the configuration file for Soar. Note that this will overwrite the existing config if already setup.",
	Run: func(cmd *cobra.Command, _ []string) {
		force, _ := cmd.Flags().GetBool("force")

		ok, err := exists("/bin/config.yml")
		if err != nil {
			fmt.Printf("Failed checking config file:\n%s", err)
			os.Exit(1)
		}

		if ok {
			if !force {
				fmt.Print("Existing config file located. Do you want to continue?\nNOTE: this will overwite the existing file (y/n): ")
				var res string
				fmt.Scanln(res)
				if !strings.HasPrefix(strings.ToLower(res), "y") {
					fmt.Print("\nCancelling...")
					os.Exit(0)
				}
			}

			if err := os.Remove("/bin/config.yml"); err != nil {
				fmt.Printf("Failed removing config file:\n%s", err)
				os.Exit(1)
			}

			file, err := os.Create("/bin/config.yml")
			if err != nil {
				fmt.Printf("Failed creating new config file:\n%s", err)
				os.Exit(1)
			}

			logc := LogConfig{
				StrictMode:       false,
				ShowDebug:        false,
				ShowHTTPLog:      false,
				ShowWebSocketLog: false,
				ErrorOutDir:      "/bin/logs",
			}

			soarc := &SoarConfig{
				Version: "0.0.1a",
				Application: struct {
					URL string "yaml:\"url\""
					Key string "yaml:\"key\""
				}{
					URL: "",
					Key: "",
				},
				Client: struct {
					URL string "yaml:\"url\""
					Key string "yaml:\"key\""
				}{
					URL: "",
					Key: "",
				},
				Logs: logc,
			}

			out, err := yaml.Marshal(&soarc)
			if err != nil {
				fmt.Printf("Unknown error:\n%s", err)
				os.Exit(1)
			}
			file.Write(out)
			fmt.Print("\nSuccessfully setup new Soar config!\nUse 'soar config set' to edit config options")
		}
	},
}

var configViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the current Soar configuration",
	Long:  "View the current Soar configuration",
	Run: func(cmd *cobra.Command, _ []string) {
		ok, err := exists("/bin/config.yml")
		if err != nil {
			fmt.Printf("Failed checking config file:\n%s", err)
			os.Exit(1)
		}
		if !ok {
			fmt.Print("Config file not found.")
			os.Exit(0)
		}

		file, err := os.ReadFile("/bin/config.yml")
		if err != nil {
			fmt.Printf("Failed checking config file:\n%s", err)
			os.Exit(1)
		}

		var config SoarConfig
		yaml.Unmarshal(file, &config)

		out := "Soar Config:\n"
		out += fmt.Sprintf("version:\t%v\n\n", config.Version)
		out += fmt.Sprintf("application:\n - url: %s\n - key: %s\n", config.Application.URL, config.Application.Key)
		out += fmt.Sprintf("client:\n - url: %s\n - key: %s\n", config.Application.URL, config.Application.Key)
		out += fmt.Sprintf("strict mode:\t%v\n", config.Logs.StrictMode)
		out += fmt.Sprintf("show debug:\t%v\n", config.Logs.ShowDebug)
		out += fmt.Sprintf("show HTTP log:\t%v\n", config.Logs.ShowHTTPLog)
		out += fmt.Sprintf("show websocket log:\t%v\n", config.Logs.ShowWebSocketLog)
		out += fmt.Sprintf("log output dir:\t%v\n", config.Logs.ErrorOutDir)

		fmt.Print(out)
	},
}

func init() {
	configSetupCmd.Flags().Bool("force", true, "Force overwrites the existing configuration file")
	ConfigCmd.AddCommand(configSetupCmd)
	ConfigCmd.AddCommand(configViewCmd)
}
