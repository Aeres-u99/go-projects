package main

import (
    "fmt"
    "log"
    "encoding/json"
    "math/rand"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

type Movie struct {
    ID string `json:"id"`
    Isbn string `json:"isbn"`
    Title string `json:"title"`
    Director *Director `json:"director"`

}

type Director struct {
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
}


var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    params := mux.Vars(r)
    for _ , item := range movies {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
}
    

func createMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json") 
    var movie Movie
    _ = json.NewDecoder(r.Body).Decode(&movie)
    movie.ID = strconv.Itoa(rand.Intn(100000))
    movies = append(movies, movie)
    json.NewEncoder(w).Encode(movie)
}


func updateMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json") 
    params := mux.Vars(r)
    for index , item := range movies {
        if item.ID == params["id"] {
            movies = append(movies[:index], movies[index+1:]...)
            var movie Movie
             _ = json.NewDecoder(r.Body).Decode(&movie)
            movie.ID = params["id"]
            movies = append(movies, movie)
            json.NewEncoder(w).Encode(movie)
        }
    }
}


func deleteMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    params := mux.Vars(r)
    for index, item := range movies {
        if item.ID == params["id"] {
            // nice move
            movies = append(movies[:index], movies[index+1:]...)
            break
        }
    }
}




func main() {
    r := mux.NewRouter()
    
    movies = append(movies, Movie{ID: "1", Isbn: "432887", Title: "Movie 1", Director: &Director{Firstname: "John", Lastname: "Cena"}})
    movies = append(movies, Movie{ID: "2", Isbn: "432888", Title: "Movie 2", Director: &Director{Firstname: "ohn", Lastname: "Cena"}})
    movies = append(movies, Movie{ID: "3", Isbn: "432889", Title: "Movie 3", Director: &Director{Firstname: "hn", Lastname: "Ca"}})
    r.HandleFunc("/movies", getMovies).Methods("GET")
    r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
    r.HandleFunc("/movies", createMovie).Methods("POST")
    r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
    r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
    
    fmt.Printf("Starting server at port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
