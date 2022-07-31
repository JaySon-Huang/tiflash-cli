package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	cobra.EnableCommandSorting = false

	rootCmd := &cobra.Command{
		Use:   "tiflash-ctl",
		Short: "TiFlash Controller",
		Long:  "TiFlash Controller (tiflash-ctl) is a command line tool for TiFlash Server",
	}
	rootCmd.AddCommand(newDispatchCmd(), newCheckCmd(), newPgCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
