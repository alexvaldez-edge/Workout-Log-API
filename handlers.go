package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// getPeople

func getPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("GET People handler")

	w.Header().Set("Content-Type", "application/json")
	var people []Person

	rows, err := db.Query("SELECT * FROM Person;")
	if err != nil {
		log.Fatal("Could not query from Person table")
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var person Person
		err := rows.Scan(&person.PersonID, &person.PersonName)
		if err != nil {
			panic(err.Error())
		}
		people = append(people, person)
	}

	json.NewEncoder(w).Encode(people)
}

// Create Person
func createPerson(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO Person(PersonName) VALUES(?)")
	if err != nil {
		log.Println("Error: Person insert not created.")
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	// PersonID := keyVal["PersonID"]
	PersonName := keyVal["PersonName"]

	_, err = stmt.Exec(PersonName)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "New Person Created")
}

func getWorkouts(w http.ResponseWriter, r *http.Request) {
	log.Println("GET Workouts handler")

	w.Header().Set("Content-Type", "application/json")
	var workouts []Workout

	rows, err := db.Query("SELECT * FROM Workout;")
	if err != nil {
		log.Fatal("Could not query from Workout table")
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var workout Workout
		err := rows.Scan(&workout.WorkoutID, &workout.WorkoutName, &workout.Description)
		if err != nil {
			panic(err.Error())
		}
		workouts = append(workouts, workout)
	}

	json.NewEncoder(w).Encode(workouts)
}

func getExercises(w http.ResponseWriter, r *http.Request) {
	log.Println("GET Exercises handler")

	w.Header().Set("Content-Type", "application/json")
	var exercises []Exercises

	rows, err := db.Query("SELECT * FROM Exercises;")
	if err != nil {
		log.Fatal("Could not query from Exercises table")
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var exercise Exercises
		err := rows.Scan(&exercise.ExercisesID, &exercise.ExerciseName, &exercise.Description)
		if err != nil {
			panic(err.Error())
		}
		exercises = append(exercises, exercise)
	}

	json.NewEncoder(w).Encode(exercises)
}
