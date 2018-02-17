package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func removePunctuation(sentence string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9\\s]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(sentence, "")
}

func longestWord(sentence string) string {
	longest := ""
	words := strings.Split(removePunctuation(sentence), " ")
	for _, word := range words {
		if len(longest) < len(word) {
			longest = word
		}
	}
	return longest
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter sentence:")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(longestWord(text))
}
