package main

import (
	"fmt"
	"strings"

	"github.com/thedevsaddam/dictionary/data"
)

func printer(keyword string, entries []data.Entry) {
	if len(entries) == 0 {
		fmt.Printf("No thing found by keyword [%s]\n", keyword)
		return
	}
	for _, e := range entries {
		definition := strings.Replace(e.Definition, keyword, "\033[1m"+keyword+"\033[0m", -1)
		if e.Wordtype == "" {
			fmt.Printf("%s:  %s\n\n", e.Word, definition)
			continue
		}
		fmt.Printf("%s (%s):  %s\n\n", e.Word, strings.TrimSpace(e.Wordtype), definition)
	}
	fmt.Println("")
}
