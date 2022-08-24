package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

func (eb englishBot) getGreeting() string {
	return "hello there"
}

type spanishBot struct{}

func (es spanishBot) getGreeting() string {
	return "hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	es := spanishBot{}

	printGreeting(eb)
	printGreeting(es)

	dog := robotDog3000{}
	dog.on()
	dog.bark()
	dog.off()
}

type machine interface {
	on()
	off()
}

type dog interface {
	bark()
}

// we can compose interfaces to make other interface
type robotDog interface {
	machine
	dog
}

type robotDog3000 struct{}

func (r robotDog3000) on() {
	fmt.Println("turning on...")
}

func (r robotDog3000) off() {
	fmt.Println("turning off...")
}

func (r robotDog3000) bark() {
	fmt.Println("Blaff, Blaff, Blaff")
}
