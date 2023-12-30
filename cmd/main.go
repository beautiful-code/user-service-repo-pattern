package main

import (
	"log"
	"net/http"
	"user-service-repo-pattern/pkg/user"
)

func main() {
	db := user.ConnectDB()
	userRepository := user.NewUserRepository(db)
	userService := user.NewService(userRepository)

	http.HandleFunc("/create", userService.CreateUserHandler)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
