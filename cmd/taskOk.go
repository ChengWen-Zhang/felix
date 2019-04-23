package cmd

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// taskdnCmd represents the taskdn command
var taskdnCmd = &cobra.Command{
	Use:   "taskok",
	Short: "set a row done in TaskList",
	Long:  `usage:felix taskok 1`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Fatal("ID must be an integer")
		}
		if err := models.TaskUpdate(uint(id), "DONE"); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(taskdnCmd)
}
