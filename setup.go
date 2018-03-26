package main

import (
	"github.com/spf13/cobra"
	"github.com/thedevsaddam/dictionary/data"
)

var setupCMD = &cobra.Command{
	Use:   "setup",
	Short: "Setup the dictionary database",
	Long:  `Setup the dictionary database`,
	Run:   setup,
}

func setup(cmd *cobra.Command, args []string) {
	data.Setup()
}
