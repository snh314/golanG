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

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// getMovies handles GET requests and returns the list of all movies in JSON format.
func getMovies(w http.ResponseWriter, r *http.Request) {
	// w: Used to SEND data/response back to the client (Postman/Frontend)
	// r: Used to RECEIVE data/request info coming from the client
	w.Header().Set("Content-Type", "Application/json") // 1. Response type set & Tell the client (Browser/Postman) to expect JSON data
	json.NewEncoder(w).Encode(movies)                  // 2. movies slice → JSON → Browser. Encode the 'movies' slice into JSON and write it directly to the response stream

}

// deleteMovie removes a specific movie from the slice based on the ID provided in the URL.
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// 1. Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")
	/* 2. Extract path variables(id) from the url (e.g., gets "2" from /movies/2). mux (Gorilla Router) extracts {id} part from URL.
	r = Browser's request, inside it URL "/movies/2" is stored
	mux.Vars(r) = extracts {id}="2" from "/movies/2" and makes a map
	params = map{"id": "2"} */
	params := mux.Vars(r)
	// 3. Iterate through the movies slice to find the matching ID
	for index, item := range movies {
		// 4. when param[id] == movies.Id; Remove the movie from the slice
		if item.Id == params["id"] {
			//  Append elements BEFORE the target index with elements AFTER the target index
			movies = append(movies[:index], movies[index+1:]...)

			// Stop searching once the item is found and deleted
			break
		}
	}
	// 5. Return the updated list of movies back to the client
	json.NewEncoder(w).Encode(movies)
}

// Find a single movie by its ID and return it as JSON. URL PATTERN: /movies/{id}  (example: /movies/2). METHOD: GET
func getMovie(w http.ResponseWriter, r *http.Request) {
	// 1. Set response content type to JSON (Fixed typo)
	w.Header().Set("Content-Type", "application/json")
	// 2. Extract the "id" variable from the request URL , via "r"
	params := mux.Vars(r)
	// 3. Iterate through the movies slice using a blank identifier (_) since index is not needed
	for _, item := range movies {
		// 4. If a matching ID is found, encode the movie data to JSON and send it
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	// 5. If the loop finishes without executing 'return', the movie was not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Movie not found"})
}

// createMovie parses JSON data from the request body, assigns a random ID, and adds the new movie to the slice.
func createMovie(w http.ResponseWriter, r *http.Request) {
	// 1. Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// 2. Create an empty Movie struct to hold the incoming data
	var movie Movie
	// 3. Decode the incoming JSON body into the 'movie' variable
	// Note: In production, you should handle the error instead of using '_'
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// 4. Generate a random integer, convert it to a string, and assign it as the new movie's ID
	movie.Id = strconv.Itoa(rand.Intn(1000000))
	// 5. Add the newly created movie to the global movies slice
	movies = append(movies, movie)
	// 6. Return the newly created movie (including its generated ID) back to the client
	json.NewEncoder(w).Encode(movie)
}

// updateMovie updates an existing movie by first deleting the old record
// and then appending the new record with the same ID.
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// 1. Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// 2. Extract the target movie ID from the URL parameters
	params := mux.Vars(r)
	// 3. Loop through the movies slice to find the matching ID
	for index, item := range movies {

		if item.Id == params["id"] {
			// 4. DELETE: Remove the old movie data from the slice
			movies = append(movies[:index], movies[index+1:]...)
			// 5. DECODE: Create a new variable and read the updated data from request body
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			// 6. ASSIGN ID: Ensure the updated movie retains the original ID from the URL
			movie.Id = params["id"]
			// 7. ADD: Append the newly updated movie back into the slice
			movies = append(movies, movie)
			// 8. Send the updated movie back as a response
			json.NewEncoder(w).Encode(movie)
			// 9. CRITICAL: Exit the function immediately so the loop doesn't continue
			return
		}
	}
}

func main() {
	movies = append(movies, Movie{Id: "1", Isbn: "12300", Title: "Movie_No_1", Director: &Director{Firstname: "Jon", Lastname: "Doe"}})
	movies = append(movies, Movie{Id: "2", Isbn: "123001", Title: "Movie_No_2", Director: &Director{Firstname: "Jenny", Lastname: "Doe"}})

	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
