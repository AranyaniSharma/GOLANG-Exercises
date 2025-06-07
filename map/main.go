package main

import "fmt"

func main() {

	colors1 := map[string]string{
		"red":   "#ff000",
		"green": "#4bb40",
		"pink":  "#2vf00",
		"white": "ffffff",
	}
	fmt.Println(colors1)
	fmt.Println(colors1["pink"])

	for colors, hex := range colors1 {
		fmt.Println(colors, hex)
	}

	var colors2 map[string]string
	fmt.Println(colors2)

	colors3 := make(map[string]string)
	fmt.Println(colors3)

	colors3["white"] = "fffff"

	colors3["white"] = "ff2ff"
	fmt.Println(colors3)

	delete(colors1, "pink")
	fmt.Println(colors1)
	printMap(colors1)

}
func printMap(c map[string]string) {
	for colors, hex := range c {
		fmt.Println(colors, hex)
	}
}
