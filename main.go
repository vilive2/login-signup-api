package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const port = "3000"

type User struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Profilepic string `json:"profilepic"`
	Password   string `json:"password"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("%+v", user)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "logged in")
	} else {
		http.Error(w, "", http.StatusNotFound)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("%+v", user)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "registered")
	} else {
		http.Error(w, "", http.StatusNotFound)
	}
}

func profile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		var user = User{
			Name:     "",
			Username: "",
			Email:    "",
		}

		w.WriteHeader(http.StatusOK)
		body, _ := json.Marshal(user)
		w.Write(body)
	} else {
		http.Error(w, "", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/profile", profile)

	fmt.Println("Server is running on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
