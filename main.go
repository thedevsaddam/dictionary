package main

import (
	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:   "dictionary",
		Short: "Offline terminal dictionary",
		Long:  `Offline terminal dictionary`,
	}
)

func init() {
	// register commands
	cmd.AddCommand(findCMD)
	cmd.AddCommand(setupCMD)
	cmd.AddCommand(fuzzyCMD)
	cmd.AddCommand(versionCmd)
}

func main() {
	cmd.Execute()
}
