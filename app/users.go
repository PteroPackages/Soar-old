package app

import (
	"fmt"

	"github.com/pteropackages/soar/request"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var GetUsersCmd = &cobra.Command{
	Use:   "get-users",
	Short: "Gets a list of users from the panel",
	Long:  "Gets a list of users from the panel, following the Pterodactyl pagination system",
	Run: func(cmd *cobra.Command, args []string) {
		session := request.NewSession("application")
		// if len(args) > 0 {
		// 	// TODO
		// }
		data, err := session.Request("/api/application/users", "GET", nil)
		if err != nil {
			panic(err)
		}

		var parsed map[string]string
		yaml.Unmarshal(data, &parsed)
		fmt.Printf("%v\n", parsed)
	},
}

func init() {
	flags := GetUsersCmd.Flags()
	flags.String("id", "", "the internal ID of the user")
	flags.String("uuid", "", "the UUID of the user")
	flags.String("username", "", "the username to query by")
	flags.String("email", "", "the email to query by")
	flags.String("external", "", "the external ID of the user")
	flags.Bool("servers", true, "whether to include servers")
	flags.Bool("s", true, "")
}
