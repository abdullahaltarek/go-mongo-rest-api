package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
}

var db *mgo.Database

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []Movie
	err := db.C("movies").Find(nil).All(&movies)

	if err != nil {
		JsonResponse(w, 501, map[string]string{"msg": "an error occurred"})
		return
	}
	JsonResponse(w, 200, movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/movies/delete/"):]
	var movie Movie
	err := db.C("movies").FindId(bson.ObjectIdHex(id)).One(&movie)
	err = db.C("movies").Remove(movie)
	if err != nil {
		JsonResponse(w, 501, map[string]string{"msg": "an error occurred"})
		return
	}
	JsonResponse(w, 200, map[string]string{"msg": "deleted successfully"})
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/movies/update/"):]
	var movie Movie
	var movieData Movie
	err := db.C("movies").FindId(bson.ObjectIdHex(id)).One(&movie)
	json.NewDecoder(r.Body).Decode(&movieData)
	fmt.Println(movieData)
	movieData.ID = movie.ID
	err = db.C("movies").UpdateId(movie.ID, movieData)
	if err != nil {
		JsonResponse(w, 501, map[string]string{"msg": "an error occurred"})
		return
	}
	JsonResponse(w, 200, map[string]string{"msg": "updated successfully"})
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	movie.ID = bson.NewObjectId()
	json.NewDecoder(r.Body).Decode(&movie)
	err := db.C("movies").Insert(movie)

	if err != nil {
		JsonResponse(w, 501, map[string]string{"msg": "an error occurred"})
		return
	}
	JsonResponse(w, 200, map[string]string{"msg": "created successfully"})
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/movies/get/"):]
	var movie Movie
	err := db.C("movies").FindId(bson.ObjectIdHex(id)).One(&movie)
	if err != nil {
		JsonResponse(w, 501, map[string]string{"msg": "an error occurred"})
		return
	}
	JsonResponse(w, 200, movie)
}

func JsonResponse(w http.ResponseWriter, code int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/movies", GetAllMovies)
	r.HandleFunc("/movies/get/", GetMovie)
	r.HandleFunc("/movies/create/", CreateMovie)
	r.HandleFunc("/movies/update/", UpdateMovie)
	r.HandleFunc("/movies/delete/", DeleteMovie)

	http.ListenAndServe(":8076", r)
}

func init() {
	conn, err := mgo.Dial("mongodb")
	if err != nil {
		panic(err.Error())
		panic("Database connection error")
	}
	db = conn.DB("moviesdb")
}
