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
	r.HandleFunc("/v1/wol/persons/{id}", updatePerson).Methods("PUT")

	r.HandleFunc("/v1/wol/workouts/name/{id}", updateWorkoutName).Methods("PUT")
	r.HandleFunc("/v1/wol/workouts/description/{id}", updateWorkoutDescription).Methods("PUT")

	r.HandleFunc("/v1/wol/exercises/name/{id}", updateExerciseName).Methods("PUT")
	r.HandleFunc("/v1/wol/exercises/description/{id}", updateExerciseDescription).Methods("PUT")

	// D --> Delete
	r.HandleFunc("/v1/wol/persons/{id}", deletePerson).Methods("DELETE")
	// r.HandleFunc("/v1/wol/workouts/{id}", deleteWorkout).Methods("DELETE")
	// r.HandleFunc("/v1/wol/exercises/{id}", deleteExercise).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
