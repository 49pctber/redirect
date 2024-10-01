package cmd

import (
	"log"

	"github.com/49pctber/redirect"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a redirect",
	Long:  `delete a redirect`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := redirect.DeleteRedirect(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
