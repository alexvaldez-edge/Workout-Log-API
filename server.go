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
	// C --> Create
	r.HandleFunc("/v1/wol/persons", createPerson).Methods("POST")
	r.HandleFunc("/v1/wol/workouts", createWorkout).Methods("POST")
	r.HandleFunc("/v1/wol/exercises", createExercise).Methods("POST")

	// R --> Read
	r.HandleFunc("/v1/wol/persons", getPeople).Methods("GET")
	r.HandleFunc("/v1/wol/workouts", getWorkouts).Methods("GET")
	r.HandleFunc("/v1/wol/exercises", getExercises).Methods("GET")

	// U --> Update
	r.HandleFunc("/v1/wol/persons", updatePerson).Methods("PUT")
	r.HandleFunc("/v1/wol/workouts", updateWorkout).Methods("PUT")
	r.HandleFunc("/v1/wol/exercises", updateExercise).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
