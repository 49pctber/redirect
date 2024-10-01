package cmd

import (
	"log"

	"github.com/49pctber/redirect"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add or update a redirect",
	Long:  `add or update a redirect`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := redirect.NewRedirect(args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		err = r.Save()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
