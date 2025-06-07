package main

import (
	"fmt"
	"time"
)

func main() {
	//// Set up channel on which to send signal notifications.
	//// We must use a buffered channel or risk missing the signal
	//// if we're not ready to receive when the signal is sent.
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//
	//// Block until a signal is received.
	//s := <-c
	//fmt.Println("Got signal:", s)

	fmt.Println(time.Now())
	fmt.Println(time.Now().String())
	myDateString := time.Now().String()
	//myDateString := "2018-01-20 04:35"
	s := myDateString[0:10]
	s1 := myDateString[0:19]
	fmt.Println(s)
	//for i:=0;i<19;i++ {
	//	myDateString=myDateString
	//}// Parse the date string into Go's time object
	// The 1st param specifies the format,
	// 2nd is our date string
	myDate, err := time.Parse("2006-01-02", s)
	//	myDate2, err := time.Parse("2006-01-02", "2022-08-19")
	if err != nil {
		panic(err)
	}
	fmt.Println(myDate)
	myDate2, err := time.Parse("2006-01-02 15:04:05", s1)
	if err != nil {
		panic(err)
	}
	fmt.Println(myDate2)
	fmt.Println(myDate2.Add(8 * time.Hour))

	if myDate2.Before(myDate2.Add(8 * time.Hour)) {
		fmt.Println("yes")
	}
	//fmt.Println(myDate.After(myDate.AddDate(0, 0, 1)))
	//if myDate.Equal(myDate2) {
	//
	//}
	//myDate.Add()
	//fmt.Println(myDate)
	//fmt.Println(myDate.Equal(myDate2))

	println(1 + 2*3)
	loginTime := time.Now().String()

	fmt.Println(loginTime)

}
