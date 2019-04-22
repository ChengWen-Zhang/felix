package cmd

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// proxyCmd represents the proxy command
var proxySocksCmd = &cobra.Command{
	Use:   "sshsocks",
	Short: "start a socks4/5 proxy",
	Long:  `usage: felix proxy socks 2 --l=1080`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			color.Red("select a ssh by ID")
			flx.AllMachines("")
			return
		}
		dbId, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Fatal("ID must be an integer", err)
		}
		h, err := models.MachineFind(uint(dbId))
		if err != nil {
			log.Fatal("ssh info is not found:", err)
		}
		if err := flx.RunSocksProxy(h, localPort); err != nil {
			log.Fatal(err)
		}
	},
}
var localPort int

func init() {
	rootCmd.AddCommand(proxySocksCmd)
	proxySocksCmd.Flags().IntVarP(&localPort, "localPort", "l", 1080, "socks4/5 local port")
	proxySocksCmd.MarkFlagRequired("localPort")
}
