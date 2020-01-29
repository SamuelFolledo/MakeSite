package main

import (
	"Io/ioutil"
	"fmt"
)

func main() {
	line, err := readFile("first-post.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Print(line)
	}
}

func readFile(fileName string) (line string, errorMessage error) { //method that will read a file and return lines or error
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		errorMessage = err
		line = ""
	} else {
		errorMessage = nil
		line = string(fileContents)
	}
	// fmt.Print(string(fileContents))
	return
}
