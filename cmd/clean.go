package cmd

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "purge all felix configuration",
	Long:  `purge all felix info by destroying SQLite database file`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := models.FlushSqliteDb(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
