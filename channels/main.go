package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// the channel is the only way to a go routine to comunicate with other
	c := make(chan string)
	// to send data to an channel
	// channel <- value
	// to receive data from the channel, is a blocking operation
	// var <- channel
	// funcX(<- channel)

	for _, link := range links {
		// the go keyword create a new go routine
		// there is a scadale to put go routines to run os the processor
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		//fmt.Println(link, "might be down")
		c <- link + " might be down"
		return
	}

	//fmt.Println(link, "is up")
	c <- link + " is up"
}
