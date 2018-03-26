package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thedevsaddam/dictionary/data"
)

var findCMD = &cobra.Command{
	Use:   "find",
	Short: "Find search words with exact match",
	Long:  `Find search words with exact match`,
	Run:   find,
}

func find(cmd *cobra.Command, args []string) {
	d := data.New()
	defer d.Close()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Dictionary Search: ")

	for {
		line, _, _ := reader.ReadLine()
		cls()
		keyword := string(line)

		fmt.Println("Dictionary Search: ", keyword)
		entries := d.Find(keyword)
		printer(keyword, entries)
	}
}
