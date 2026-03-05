package main

import (
		"fmt"
		"os"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func create(){
	new, err := os.Create("new.txt")

	check(err)
	fmt.Println(new)

	fold, err := new.WriteString("This is Our new String")
	check(err)

	fmt.Println(fold)
}


func read(){
	file, err := os.ReadFile("sample.txt")

	check(err)

	fmt.Println(file) // This gives us a slice of bytes of the content of the text


	fmt.Println(string(file)) // This gives us the content of the file because we convert the slice of bytes into string.
}

func main(){
	create()
	read()
}