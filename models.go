package main

// Person Struct (Model)
type Person struct {
	PersonID   uint64 `json:"PersonID"`
	PersonName string `json:"PersonName"`
}

// Workout Struct (Model)
type Workout struct {
	WorkoutID   uint64 `json:"WorkoutID"`
	WorkoutName string `json:"WorkoutName"`
	Description string `json:"Description"`
}

// Exercises Struct (Model)
type Exercises struct {
	ExercisesID  uint64 `json:"ExercisesID"`
	ExerciseName string `json:"ExerciseName"`
	Description  string `json:"Description"`
}
