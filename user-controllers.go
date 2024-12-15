package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

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
			ErrorJSON(w, err, http.StatusOK)
			return
		}

		err = UserServices{}.Login(user)
		if err != nil {
			ErrorJSON(w, err, http.StatusOK)
			return
		}

		user.Password = ""
		WriteJSON(w, user)
	} else {
		ErrorJSON(w, errors.New("method not found"), http.StatusNotFound)
	}
}

func register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodPost {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			ErrorJSON(w, err, http.StatusOK)
			return
		}
		defer r.Body.Close()

		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			ErrorJSON(w, err, http.StatusOK)
			return
		}

		err = UserServices{}.Register(user)
		if err != nil {
			ErrorJSON(w, err, http.StatusOK)
			return
		}

		WriteJSON(w, nil)
	} else {
		ErrorJSON(w, errors.New("method not found"), http.StatusNotFound)
	}
}

func profile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		username := r.URL.Query().Get("username")
		if username == "" {
			ErrorJSON(w, errors.New("query param username missing"), http.StatusOK)
			return
		}

		user, err := UserServices{}.Profile(username)
		if err != nil {
			ErrorJSON(w, err, http.StatusOK)
			return
		}

		WriteJSON(w, user)
	} else {
		ErrorJSON(w, errors.New("method not found"), http.StatusNotFound)
	}
}
