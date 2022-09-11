package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	//http.HandleFunc("/", getRoot)
	//http.HandleFunc("/hello", getHello)
	//
	//err := http.ListenAndServe(":3333", nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/create", getBody)
	mux.HandleFunc("/form", getForm)
	mux.HandleFunc("/echo/json", echoJson)

	err := http.ListenAndServe(":3333", mux)

	if err != nil {
		fmt.Println("Error when starting server")
		fmt.Println(err)
		os.Exit(1)
	}
}

// curl 'http://localhost:3333/?first=1&second=2'
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")
	fmt.Printf("hasFirst: %t, value: %s\n", hasFirst, first)
	fmt.Printf("hasSecond: %t, value: %s\n", hasSecond, second)
	io.WriteString(w, "This is my web site")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /hello request")
	io.WriteString(w, "Hello HTTP")
}

// curl -X POST -d '{"name":"fulano","age":20}' http://localhost:3333/create
func getBody(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error when reading the request body")
		fmt.Println(err)
	}

	fmt.Printf("%s\n", body)
}

// curl -v -X POST -F 'name=fulano' http://localhost:3333/form
func getForm(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	if name == "" {
		w.Header().Set("x-missing-field", "name")
		// once this call is done the header can't be changed
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(name)
	io.WriteString(w, fmt.Sprintf("Hello, %s", name))
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// curl -X POST -d '{"name":"fulano","age":20}' http://localhost:3333/create
func echoJson(w http.ResponseWriter, r *http.Request) {
	var jsonBody User
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error when reading body")
	}
	fmt.Printf("%s\n", body)

	err = json.Unmarshal(body, &jsonBody)

	if err != nil {
		fmt.Println("Error when unmarshal")
	}
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(jsonBody)

	if err != nil {
		fmt.Println("Error when encoding")
	}
}
