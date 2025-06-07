package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "Hello this is file which I am creating"

	file, err := os.Create("myfile")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println("lenghth is ", length)
	defer file.Close()

	//another way
	err2 := ioutil.WriteFile("myfile2", []byte(content), 0666)
	if err2 != nil {
		fmt.Println("Error2 is", err2)
		os.Exit(1)
	}
	content2, err3 := ioutil.ReadFile("myfile2")
	if err3 != nil {
		//option 1 log the error and return a newDeck())
		//option 2 log the error and entirely quit the program

		panic(err)
	}
	fmt.Println(len(content2))
	fmt.Println(string(content2))

}
