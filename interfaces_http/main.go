package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	//bs := make([]byte, 99999) //empty bite slice which has 99999 empty spaces

	//fmt.Println(bs)
	//resp.Body.Read(bs) //by this all the html will be read to byte slice
	//fmt.Println(string(bs))
	//
	//content, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Error", err)
	//	os.Exit(1)
	//}
	//fmt.Println("Here is the content ", string(content))

	io.Copy(os.Stdout, resp.Body) //writer interface,reader interface

	fmt.Println("now logwriter")
	lw := logWriter{}
	io.Copy(lw, resp.Body) //as writer interface has function write so all the types which can access write are suitable here

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote", len(bs), "slices")
	return len(bs), nil
}
