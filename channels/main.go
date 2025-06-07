package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//do not communicate by sharing memory instead share memory by communicating
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}
	//for _ = range links {
	//
	//	fmt.Println(<-c)
	//}
	//infinte times check
	//for {
	//	go checkLink(<-c, c)
	//}
	//for l := range c {
	//	go checkLink(l, c)
	//}

	for l := range c {
		var li string = l
		go func(li string) {
			time.Sleep(5 * time.Second)
			checkLink(li, c)
		}(li)

	}
}

func checkLink(link string, c chan string) {

	//fmt.Println(&link)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!", err)
		//c <- "might be down"
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	//c <- "yep it's up"
	c <- link
}
