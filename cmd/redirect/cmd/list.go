package cmd

import (
	"fmt"
	"log"

	"github.com/49pctber/redirect"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all redirects",
	Long:  `list all redirects`,
	Run: func(cmd *cobra.Command, args []string) {
		rs, err := redirect.GetAllRedirects()
		if err != nil {
			log.Fatal(err)
		}

		for _, r := range rs {
			fmt.Println(r)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
