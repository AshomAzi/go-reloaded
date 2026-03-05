package main

import (
	"fmt"
	"os"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func two(file1, file2 string){
	f1, err := os.ReadFile(file1)
	check(err)
	fmt.Println(string(f1))


	f2, err := os.ReadFile(file2)
	check(err)
	fmt.Println(string(f2))
}

func main(){
	two("sample.txt", "new.txt")
}