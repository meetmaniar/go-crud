package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Job struct {
	JobId string `json:"jobId"`
	Job   string `json:"job"` // the convention is key:value and not key: value -> `json:"job"` and not `json: "job"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Job  *Job   `json:"job"`
}

func createUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)
	//rand.Intn(1000) will create a random number between 1 and 1000
	user.ID = strconv.Itoa(rand.Intn(10000))
	users = append(users, user)
	json.NewEncoder(response).Encode(user)
}

func updateUser(response http.ResponseWriter, request *http.Request) {
	// Set JSON Conent Type
	response.Header().Set("Content-Type", "application/json")
	// Getting the params from the http request
	var updatedUser User
	_ = json.NewDecoder(request.Body).Decode(&updatedUser)
	params := mux.Vars(request)
	for i := 0; i < len(users); i++ {
		if params["id"] == users[i].ID {
			updatedUserInfo(&users[i], updatedUser)
			json.NewEncoder(response).Encode(users[i])
			return
		}
	}
}

func updatedUserInfo(user *User, updatedUser User) {
	// user.ID = user.ID // We don't need this call because we don't want to update the id
	user.Name = updatedUser.Name
	user.Job = updatedUser.Job
	fmt.Println(user)
}

func getUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(users)
}

func getUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for i := 0; i < len(users); i++ {
		if params["id"] == users[i].ID {
			json.NewEncoder(response).Encode(users[i])
			return
		}
	}
}

func deleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	// index=0
	// for (item : users) {}
	fmt.Printf("param:" + params["id"])
	for index, item := range users {
		if item.ID == params["id"] {
			//The below line will take the element with the passed parameter and append the rest of the elements.
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	// Returning the rest of the movies
	json.NewEncoder(response).Encode(users)
}

var users []User

func main() {
	users = append(users, User{ID: "user_1", Name: "John Doe", Job: &Job{JobId: "job_1", Job: "Software Engineer"}})
	users = append(users, User{ID: "user_2", Name: "Harry Potter", Job: &Job{JobId: "job_2", Job: "Wizard"}})
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/createUser", createUser).Methods("POST")
	r.HandleFunc("/updateUser/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/deleteUser/{id}", deleteUser).Methods("DELETE")

	fmt.Printf("Starting Server at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
