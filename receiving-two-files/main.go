package main

import (
	"fmt"
	"os"
)

func check(e error){
	if e != nil {
		fmt.Println("An Error occured")
		os.Exit(1)
	}
}

func two(file1, file2 string){
	f1, err := os.ReadFile(file1)
	check(err)

	f1_result := string(f1)

	f2, err := os.Create(file2)
	check(err)
	fmt.Println("Copied Successfully at the address:", f2)

	f2_result, err := f2.WriteString(f1_result)
	check(err)

	fmt.Println(f2_result)
}

func main(){
	two("sample.txt", "new.txt")
}