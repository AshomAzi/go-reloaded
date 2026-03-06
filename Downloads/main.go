package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// file, err := os.Open("sample.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// // Outer Scanner: Reads the file line by line
	// lineScanner := bufio.NewScanner(file)
	// for lineScanner.Scan() {
	// 	line := lineScanner.Text()

	// 	// Inner Scanner: Reads the current line word by word
	// 	wordScanner := bufio.NewScanner(strings.NewReader(line))
	// 	wordScanner.Split(bufio.ScanWords)

	// 	var processedWords []string
	// 	for wordScanner.Scan() {
	// 		word := wordScanner.Text()

	// 		// --- YOUR LOGIC HERE ---
	// 		// Example: if word == "(hex)", convert previous word to decimal
	// 		// For now, we just collect them
	// 		processedWords = append(processedWords, word)
	// 	}

	// 	// Join the words back with spaces and print as a single line
	// 	fmt.Println(strings.Join(processedWords, " "))
	// }

	// if err := lineScanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
