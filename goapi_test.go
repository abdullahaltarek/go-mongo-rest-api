package main

import (
	"testing"
	"gopkg.in/mgo.v2"
	"net/http"
	"bytes"
)

func TestDatabaseConnection(t *testing.T) {
	_, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		t.Error("Database connection error")
	}
}

func TestGetAllMovies(t *testing.T) {
	res, err := http.Get("http://127.0.0.1:8076/movies")

	if err != nil || res.StatusCode != 200 {
		t.Errorf("%s", "Error")
		t.Error(err.Error())
	}

	defer res.Body.Close()
}

func TestGetMovie(t *testing.T) {
	res, err := http.Get("http://127.0.0.1:8076/movies/get/5b1aaffd0e78a33a005e44dc")

	if err != nil || res.StatusCode != 200 {
		t.Errorf("%s", "Error")
		t.Error(err.Error())
	}

	defer res.Body.Close()
}

func TestCreateMovie(t *testing.T) {
	res, err := http.Post("http://127.0.0.1:8076/movies/create/", "application/json", bytes.NewBufferString(`{"name": "InfluxBD", "cover_image": "", "description": "This is a data driven movie"}`))

	if err != nil || res.StatusCode != 200 {
		t.Error()
	}
}