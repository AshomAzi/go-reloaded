package main

import (
	"bufio"
	"log"
	"os"
	
	"go-reloaded/modification" // Ensure this matches your go.mod name
)

func main() {
	// 1. Open Source File
	f1, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	// 2. Create Destination File
	f2, err := os.Create("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	// 3. Scan and Delegate Logic
	scanner := bufio.NewScanner(f1)
	writer := bufio.NewWriter(f2)

	for scanner.Scan() {
		line := scanner.Text()
		
		// Call our subfolder function
		modifiedLine := modification.ProcessLine(line)
	// 2. Then handle casing (up/low)
		modifiedLine = modification.ConvertCase(modifiedLine)
		
		modifiedLine = modification.Punctuation(modifiedLine)
	
		modifiedLine = modification.FixArticles(modifiedLine)		

		_, err := writer.WriteString(modifiedLine + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Important: Flush the buffer to ensure all data is written to the file
	writer.Flush()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}