package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`              //password will not be displayed
	Tags     []string `json:"tags,omitempty"` //omitempty means if you wants the field ehich do not hve value must not b shown
}

func main() {
	EncodeJson()
}
func EncodeJson() {
	courses := []course{
		{Name: "react js", Price: 100, Password: "abc123", Platform: "loc", Tags: []string{"web", "js"}},
		{Name: "node js", Price: 230, Password: "def1234", Platform: "loc", Tags: []string{"web", "js"}},
		{Name: "angular js", Price: 400, Password: "angular@123", Platform: "loc"},
	}
	fmt.Println(courses)

	//package this data as json data

	finalJson, err := json.Marshal(courses)
	fmt.Println(string(finalJson))
	fmt.Printf("%s", finalJson)
	if err != nil {
		panic(err)
	}
	finalIndentJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalIndentJson))

}
