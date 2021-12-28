package app

import (
	"github.com/spf13/cobra"
)

var GroupCommand = &cobra.Command{
	Use: "app",
}

func init() {
	GroupCommand.AddCommand(getUsersCmd)
}
