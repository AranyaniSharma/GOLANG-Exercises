package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	PerformGetRequest()
	PerformPostRequest()
	PerformFormRequest()
}
func PerformGetRequest() {
	const myurl = "http://localhost:8000"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status code", response.StatusCode)
	fmt.Println("Content Length", response.ContentLength)

	//contentInByte, _ := ioutil.ReadAll(response.Body)
	//fmt.Println("content from read All ", string(contentInByte))

	lengthOfContent, _ := io.Copy(os.Stdout, response.Body)
	fmt.Println()
	fmt.Println("lengthOfContent is", lengthOfContent)
	fmt.Printf("tyep ofcontent is %T", lengthOfContent)
	fmt.Println()
	defer response.Body.Close()

}
func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"
	requestBody := strings.NewReader(`
	{
		"courseName":"Let's Connect ",
		"price":"0",
		"platform":"learncodeonline.in"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		fmt.Println("Eroor is   ", err)
		os.Exit(1)
	}
	contentInByte, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Post Request,\n", string(contentInByte))
	defer response.Body.Close()
}
func PerformFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata
	s1 := []string{"a", "b"}
	s2 := []string{"hello", "world"}
	data := url.Values{
		"A":       s1,
		"B":       s2,
		"Capital": []string{"A", "B", "C"},
	}
	data.Add("firstName", "Aranyani")
	data.Add("LastName", "Sharma")
	//fmt.Println(data)

	response, err := http.PostForm(myurl, data)
	if err != nil {
		fmt.Println("Eroor is   ", err)
		os.Exit(1)
	}
	contentInByte, _ := ioutil.ReadAll(response.Body)
	fmt.Println("PostForm Request,\n", string(contentInByte))
	defer response.Body.Close()
}
