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


func create(){
	file, err := os.Create("sample.txt") // Creates a new or truncates a file and returns the address to the file in memory in the variable file
	check(err)
	
	fmt.Println("File Creater successfully!",file) // Returns the address of the file

	new, err := file.WriteString("Welcome, this is a new string!") // Writes a new content to the file and returns the total number of content in the file in the variable new
	check(err)

	fmt.Println(new) // This returns the total number of itens in the file
}

func main(){
	create()
}