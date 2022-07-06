package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var fileServer = http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	const PORT = "8080"
	fmt.Println("Staring the serve on port", PORT, "...")

	if error := http.ListenAndServe(":"+PORT, nil); error != nil {
		log.Fatal(error)
	}
}

func formHandler(responseWriter http.ResponseWriter, request *http.Request) {

	if error := request.ParseForm(); error != nil {
		fmt.Fprintf(responseWriter, "ParseForm error: %v", error)
	}

	fmt.Fprintf(responseWriter, "POST request successful")

	var name = request.FormValue("name")
	var address = request.FormValue("address")

	fmt.Fprintf(responseWriter, "Name = %s", name)
	fmt.Fprintf(responseWriter, "Address = %s", address)
}

func helloHandler(responseWriter http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/hello" {
		http.Error(responseWriter, "404 not found", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(responseWriter, "Method os not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(responseWriter, "hello")
}
