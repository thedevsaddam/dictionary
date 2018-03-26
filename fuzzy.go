package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thedevsaddam/dictionary/data"
)

var fuzzyCMD = &cobra.Command{
	Use:   "fuzzy",
	Short: "Fuzzy search words with any match",
	Long:  `Fuzzy search words with any match`,
	Run:   fuzzy,
}

func fuzzy(cmd *cobra.Command, args []string) {
	d := data.New()
	defer d.Close()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Dictionary Fuzzy Search: ")

	for {
		line, _, _ := reader.ReadLine()
		cls()
		keyword := string(line)
		fmt.Println("Dictionary Fuzzy Search: ", keyword)
		entries := d.Fuzzy(keyword)
		printer(keyword, entries)
	}
}
