package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go movies up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {

	rdOne, _ := time.Parse("2006-01-02", "1986-03-07")
	rdTwo, _ := time.Parse("2006-01-02", "1981-06-12")

	movies := []models.Movie{
		models.Movie{
			ID:          1,
			Title:       "Highlander",
			ReleaseDate: rdOne,
			MPAARating:  "R",
			RunTime:     116,
			Description: "Just some description text",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		models.Movie{
			ID:          2,
			Title:       "Raiders of the Lost Ark",
			ReleaseDate: rdTwo,
			MPAARating:  "PG-13",
			RunTime:     115,
			Description: "Another stupid description text",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	out, err := json.Marshal(movies)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
