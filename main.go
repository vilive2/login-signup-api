package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "3000"

var userRepository UserRepository

func main() {
	userRepository.DB = GetDB()
	if userRepository.DB == nil {
		log.Fatal("failed to load DB\n")
	}

	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/profile", profile)

	fmt.Println("Server is running on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
