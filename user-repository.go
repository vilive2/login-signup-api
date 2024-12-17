package main

import "database/sql"

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Photo    string `json:"photo"`
	Password string `json:"password"`
}

type UserRepository struct {
	DB *sql.DB
}

func (ur *UserRepository) QueryByUsername(username string) (User, error) {
	var user User

	err := ur.DB.QueryRow(`
		SELECT 
			username,
			name,
			email,
			photo
		FROM
			users
		WHERE
			username = ?
	`, username).Scan(&user.Username, &user.Name, &user.Email, &user.Photo)

	return user, err
}

func (ur *UserRepository) Save(user *User) error {
	_, err := ur.DB.Exec(`
		INSERT INTO users(username, name, email, password, photo)
			VALUES (?, ?, ?, ?, ?);
	`, user.Username, user.Name, user.Email, user.Password, user.Photo)

	return err
}

func (ur *UserRepository) QueryUserPassword(username string) (string, error) {
	var hash string

	err := ur.DB.QueryRow(`
		SELECT 
			password
		FROM
			users
		WHERE
			username = ?
	`, username).Scan(&hash)

	return hash, err
}
