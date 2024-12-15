package main

import "database/sql"

type User struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Profilepic string `json:"profilepic"`
	Password   string `json:"password"`
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
			email
		FROM
			users
		WHERE
			username = ?
	`, username).Scan(&user.Username, &user.Name, &user.Email)

	return user, err
}

func (ur *UserRepository) Save(user *User) error {
	_, err := ur.DB.Exec(`
		INSERT INTO users(username, name, email, password)
			VALUES (?, ?, ?, ?);
	`, user.Username, user.Name, user.Email, user.Password)

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
