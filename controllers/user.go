package controllers

import (
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct{}
)

func NewUserController() *UserController {
	return &UserController{}
}

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	u := models.User{
		Name:   "Bob Smith",
		Gender: "male",
		Age:    50,
		Id:     params.ByName("id"),
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, status code, payload
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	fmt.Fprintf(writer, "%s", uj)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.User{}

	// Populate the user data
	json.NewDecoder(request.Body).Decode(&u)

	// Add an Id
	u.Id = "foo"

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(201)
	fmt.Fprintf(writer, "%s", uj)
}


// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// TODO: only write status for now
	writer.WriteHeader(200)
}

