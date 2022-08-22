package main

import "fmt"

// Go has value types and Reference types
// when we pass a value type, if we wnat to modify it we need a poiter
// when we pass a reference type, it has a reference to the original data/data struct so we can modify the original data
// values types
// int, float, string, bool and struct
// reference type
// slices, maps, channels, pointers and functions

type contacInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contacInfo
}

// go is a pass by value lang
// receiver function: the (p person) makes this function be bind with the struct person.
// the (p person) is the receiver, and p is a pointer to a copy of the original instace
// the (p *person) p is a pointer to the instace where print was called
// we can call print like
// p1.print()
// or
// pPointer := &p
// pPointer.print()
// if the receiver is a copy or an poniter go will pass the right value, a copy or reference
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func main() {
	//p := person{"Fulano", "de tal"}
	//fmt.Println(p)

	p2 := person{firstName: "foo", lastName: "bar"}
	fmt.Println(p2)

	// when creating a variable like this, the go assign the zero value for the struct feilds, not nil
	// string = ""
	// int = 0
	// float = 0
	// bool = false
	var p3 person
	fmt.Println(p3)

	p3.firstName = "cicrano"
	p3.lastName = "123"
	fmt.Printf("%+v\n", p3)

	//every feild should have a , at the end.
	p4 := person{
		firstName: "Fulano",
		lastName:  "de Tal",
		contact: contacInfo{
			email:   "fulano@mail.com",
			zipCode: 00000,
		},
	}
	p4.print()
}
