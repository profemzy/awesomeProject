package main

import (
	"awesomeProject/controllers"
	// Standard library packages
	"fmt"
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Add a handler on /test
	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Simply write some test data for now
		fmt.Fprint(w, "Welcome Home!\n")
	})

	// Get a UserController instance
	uc := controllers.NewUserController()

	// Get a user resource
	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.RemoveUser)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}
