package main

import "fmt"

type Processor struct{}

func (p Processor) process(value string) {
	fmt.Println(value)
}
