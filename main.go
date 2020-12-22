package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

// Person Struct
type Person struct {
	PersonID   int    `json:"PersonID"`
	PersonName string `json:"PersonName"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func createNewPerson(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Person Person
	json.Unmarshal(reqBody, &Person)
	// db.Create(&person)
	db.Select("PersonID", "PersonName").Create(&Person)
	fmt.Println("Endpoint Hit: Creating New Person")
	json.NewEncoder(w).Encode(&Person)
}

func returnAllPersons(w http.ResponseWriter, r *http.Request) {
	persons := []Person{}
	db.Find(&persons)
	fmt.Println("Endpoint Hit: returnAllPersons")
	json.NewEncoder(w).Encode(persons)
}

func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:8000/")
	log.Println("Quit the server with CONTROL-C.")
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/new-person", createNewPerson).Methods("POST")
	myRouter.HandleFunc("/all-persons", returnAllPersons)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := envs["DB_USER"]
	password := envs["DB_PASSWORD"]
	database := envs["DB_NAME"]

	db, err = gorm.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/"+database+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	db.AutoMigrate(&Person{})
	handleRequests()
}
