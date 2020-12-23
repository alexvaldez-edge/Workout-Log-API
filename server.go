package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server starting at http://localhost:8000/")

	connectDatabase()

	// Init mux router
	r := mux.NewRouter()

	// Handlers CRUD
	// R --> Read
	r.HandleFunc("/v1/wol/persons", getPeople).Methods("GET")
	r.HandleFunc("/v1/wol/workouts", getWorkouts).Methods("GET")
	r.HandleFunc("/v1/wol/exercises", getExercises).Methods("GET")

	// C --> Create
	r.HandleFunc("/v1/wol/persons", createPerson).Methods("POST")
	// r.HandleFunc("/v1/wol/workouts", createWorkout).Methods("POST")
	// r.HandleFunc("/v1/wol/exercises", createExercise).Methods("POST")

	// route := mux.NewRouter()

	// addAppRoutes(route)

	log.Fatal(http.ListenAndServe(":8000", r))
}
