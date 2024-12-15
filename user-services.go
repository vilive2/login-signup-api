package main

import (
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserServices struct{}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (UserServices) Login(user User) error {
	log.Printf("%+v", user)

	hash, err := userRepository.QueryUserPassword(user.Username)
	if err != nil {
		log.Println(err)
		return errors.New("server error")
	}

	if !checkPasswordHash(user.Password, hash) {
		return errors.New("invalid creds")
	}

	return nil
}

func (UserServices) Register(user User) error {
	var err error
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return errors.New("server error")
	}

	log.Printf("%+v", user)

	err = userRepository.Save(&user)
	if err != nil {
		log.Println(err.Error())
		return errors.New("server error")
	}

	return nil
}

func (UserServices) Profile(username string) (User, error) {
	user, err := userRepository.QueryByUsername(username)

	if err == sql.ErrNoRows {
		return User{}, errors.New("user not found")
	} else if err != nil {
		log.Println(err)
		return User{}, errors.New("server error")
	}

	return user, nil
}
