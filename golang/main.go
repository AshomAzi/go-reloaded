package main

import (
	"bufio"
	"golang/hexv"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// func main() {
// 	file, err := os.Open("sample.txt")
// 	check(err)
// 	defer file.Close()

// 	dst, err := os.Create("new.txt")
// 	check(err)
// 	defer dst.Close()

// 	look := bufio.NewScanner(file)
// 	for look.Scan() {
// 		line := look.Text()

// 		var modify strings.Builder

// 		for _, char := range line {
// 			upperChar := strings.ToUpper(string(char))
// 			modify.WriteString(upperChar)
// 		}

// 		_, err := dst.WriteString(modify.String() + "\n")
// 		check(err)
// 	}

// 	if err := look.Err(); err != nil {
// 		log.Fatal("Error reading file:", err)
// 	}
// }

func main() {
	file, err := os.Open("sample.txt")
	check(err)
	defer file.Close()

	dst, err := os.Create("new.txt")
	check(err)
	defer dst.Close()

	look := bufio.NewScanner(file)
	for look.Scan() {
		line := look.Text()

		// Use the package name prefix to call the function
		modifiedLine := hexv.ProcessLine(line)

		_, err := dst.WriteString(modifiedLine + "\n")
		check(err)
	}

	if err := look.Err(); err != nil {
		check(err)
	}
}
