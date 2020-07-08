package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	//splitWord := strings.Split(s, " ")
	splitWord := strings.Fields(s)
	for _, char := range splitWord {
		wordCount[char]++
	}
	return wordCount
}


func main() {
	wc.Test(WordCount)
}

