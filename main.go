package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go-gorm-rest-api/db"
	"go-gorm-rest-api/models"
	"go-gorm-rest-api/routes"
	"net/http"
)

func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{dni}", routes.DeleteUserHanlder).Methods("DELETE")
	r.HandleFunc("/user/delete", routes.GetUserDeleteHanlder).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	http.ListenAndServe(":3000", handler)
}
