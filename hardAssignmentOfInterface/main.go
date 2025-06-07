package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.Copy(os.Stdout, file)
	fmt.Println(bytes)

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("error", err)
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}
