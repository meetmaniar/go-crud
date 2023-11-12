package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Job struct {
	JobId string `json:"jobId" example:"job_1"`
	Job   string `json:"job" example:"Software Developer"` // the convention is key:value and not key: value -> `json:"job"` and not `json: "job"`
}

type User struct {
	ID   string `json:"id" example:"user_1"`
	Name string `json:"name" example:"John Doe"`
	Job  *Job   `json:"job"`
}

// createUser godoc
// @Summary POST Create new User
// @Description POST create new User
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {array} User
// @Router /createUsers [post]
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

// @title User API
// @version 1.0
// @description This is a sample service for managing users
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email nocontact@domain.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	users = append(users, User{ID: "user_1", Name: "John Doe", Job: &Job{JobId: "job_1", Job: "Software Engineer"}})
	users = append(users, User{ID: "user_2", Name: "Harry Potter", Job: &Job{JobId: "job_2", Job: "Wizard"}})
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/createUser", createUser).Methods("POST")
	router.HandleFunc("/updateUser/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/deleteUser/{id}", deleteUser).Methods("DELETE")
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	fmt.Printf("Starting Server at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
