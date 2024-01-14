package  main

import (
	"fmt"
	"log" // used to log out errors when connectin to the server
	"encoding/json" // encoding the data to json so that it is being sent to json
	"math/rand" // used in the case where we create a new movie
	"net/http"  // standard module or package for basic routing
	"strconv" // the id created from math/rand will give integers, strconv will convert them to a string
	"github.com/gorilla/mux"
)

// we will be utilizing struct and slices to create data and not a database

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`  //nesting the Director struct to the Movie struct

}

type Director struct {
	Firstname string `json:"firstname"`  // we defining the json like this so we can encode the code in json when it comes fromm postman
	Lastname string `json:"lastname"`
}


var movies []Movie // creating a varible movies for a slice named Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")  // 
	json.NewEncoder(w).Encode(movies)  // rturns all movies and converting the movies(slice data type) into a json type used in postman
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r) //This function is provided by the gorilla/mux package. It takes an http.Request object as an argument, which represents the current HTTP request being processed by the handler...... The line params := mux.Vars(r) in the context of a Go web server using the gorilla/mux package is used to extract the route variables from the URL.
	for index, item := range movies {

		if item.ID == param["id"] {
			movies = append(movies[:index], movies[index+1:]... )  // The line movies = append(movies[:index], movies[index+1:]...) in Go is used to remove an element from a slice at a specific index. Let's break it down to understand how it works:
			break
		}
	}
	json.NewEncoder(w).Encode(movies)  // return the existing movies after deleting a particular movie 

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for _, item := range movies { ////

		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item) // returns a particular movie and  converting the go valus to json
			return
		}
	}
}


func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)  /// getting the json data and  converting it to a readale format for Golang
	movie.ID = strconv.Itoa(rand.Intn(100000000))  // creating a randow id for the new movie
	movies = append(movies, movie) // appending the new movie to the existing slice of movies
	json.NewEncoder(w).Encode(movie)  // converting the movie data type to a json format 
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")
	// getting param value
	param := mux.Vars(r)
	// looping over movies using range
	for index, item := range movies {
		// deleting the movie with the id you have sent
		if item.ID == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			// adding a new movie - the movie that we send in the body of postman 
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = param["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)  // converting movie to a json format
			return
		}
	}

}



func main() {  // execution function
  
	movies = append(movies, Movie{ID: "1", Isbn: "344553", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "David"}})  // appending the struct data into a list of slice
	movies = append(movies, Movie{ID: "2", Isbn: "4848495", Title: "Movie Two", Director: &Director{Firstname: "Charles", Lastname: "Murphy"}})


	r := mux.NewRouter()  // cretating a router for the server 
	r.HandleFunc("/movies", getMovies).Methods("GET")  // router for getMovies
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET") // router for getMovie
	r.HandleFunc("/movies", createMovie).Methods("POST")  // router for createMovie
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server starting on port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))  /// creating the port the server listens on and logging an error if there is

}


