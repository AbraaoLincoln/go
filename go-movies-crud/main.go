package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const CONTENT_TYPE = "Content-type"
const APPLICATION_JON = "application/json"

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	var router = mux.NewRouter()

	addMoviedToMovieSlice()

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", creatMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	const PORT = "8000"
	fmt.Println("String the server at port", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}

func addMoviedToMovieSlice() {
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "458567", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})
}

func getMovies(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("GET movies")
	responseWriter.Header().Set(CONTENT_TYPE, APPLICATION_JON)
	json.NewEncoder(responseWriter).Encode(movies)
}

func deleteMovie(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("DELETE movie")
	responseWriter.Header().Set(CONTENT_TYPE, APPLICATION_JON)
	var params = mux.Vars(request)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("GET movie")
	responseWriter.Header().Set(CONTENT_TYPE, APPLICATION_JON)
	var params = mux.Vars(request)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(responseWriter).Encode(item)
			return
		}
	}
}

func creatMovie(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("CREATE movie")
	responseWriter.Header().Set(CONTENT_TYPE, APPLICATION_JON)
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(responseWriter).Encode(movie)
}

func updateMovie(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("UPDATE movies")
	responseWriter.Header().Set(CONTENT_TYPE, APPLICATION_JON)
	var params = mux.Vars(request)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			movie.ID = params["id"]
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(responseWriter).Encode(movie)
			return
		}
	}
}
