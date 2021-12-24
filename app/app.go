package app

import (
	"github.com/spf13/cobra"
)

var GroupCmdApplication = &cobra.Command{
	Use: "app",
}

func init() {
	GroupCmdApplication.AddCommand(GetUsersCmd)
}
