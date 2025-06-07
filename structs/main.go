package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
	fmt.Println(alex.firstName, alex.lastName)

	var garima person

	garima.firstName = "Garima"
	garima.lastName = "Sharma"

	fmt.Println(garima)
	fmt.Printf("%+v", garima)
	fmt.Println()
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jimparty@gmail.com",
			zipCode: 474003,
		},
	}
	fmt.Printf("%+v", jim)

	jim.print()
	jim.updateName("jimmy")

	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
	jim.updateName("GARIMA")
	jim.print()

	n := 2
	getNo(&n)
	wer(&jim)

	a := 2
	var p *int = &a
	fmt.Println(p, *p, &p, &a, a)

}

func (p person) print() {
	fmt.Println()
	fmt.Printf("%+v", p)

}

func (pointerToPerson *person) updateName(newFirstName string) {
	fmt.Println("comparisonss::::::::")
	fmt.Println(*pointerToPerson, pointerToPerson)
	(*pointerToPerson).firstName = newFirstName
}
func wer(pointerToPerson *person) {
	fmt.Println("comparisonss2::::::::")
	fmt.Println(*pointerToPerson, pointerToPerson)
	(*pointerToPerson).firstName = "alex"
}
func getNo(n *int) {
	fmt.Println("comparisons for number::::::::")
	fmt.Println(n, *n)
}
