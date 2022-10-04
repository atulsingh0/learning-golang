package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	// Opening a file
	fh, err := os.OpenFile("/tmp/RandomWords", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Unable to Open file:", err.Error())
	}
	defer fh.Close()

	// Generating list of lines having random words
	randomStrings(fh, readWords(), 200000, 10, 20)

}

func randomStrings(fh *os.File, wordsList []string, numLines int, minWord int, maxWord int) {

	// Total no of words in wordList
	totalWords := len(wordsList)

	for i := 0; i < numLines; i++ {
		// Picking random no of words for line
		lenWord := randomInt(minWord, maxWord)
		line := make([]string, lenWord)
		for j := 0; j < lenWord; j++ {
			// Pick and Assign a random word from wordList
			line[j] = wordsList[randomInt(0, totalWords)]
		}
		// join random words and form a line
		_, err := fh.WriteString(strings.Join(line, " ") + "\n")
		if err != nil {
			log.Fatal("Unable to write line:", err.Error())
		}
	}

}

func readWords() []string {
	// Reading wordList from file
	data, err := os.ReadFile("/usr/share/dict/words")
	if err != nil {
		log.Fatal("Unable to read data from wordList file:", err.Error())
	}
	return strings.Split(string(data), "\n")
}

func randomInt(min int, max int) int {
	// Return random interger b/w min and max
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
