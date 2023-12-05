package user

import (
	"auth/internals/utils"
	"fmt"
)

type UserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	ID        int    `json:"-"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Save() is a method defined on the `User` struct. It is used to save a user object to
// the database.
func (u *User) Save() error {
	db, err := utils.OpenDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (username, email, password) VALUES (?,?,?)")

	if err != nil {
		return err
	}
	fmt.Println(u)
	_, err = stmt.Exec(u.Username, u.Email, u.Password)

	if err != nil {
		return err
	}

	return nil
}

// GetUser retrieves a user from the database based on their username or email.
func GetUser(login string) (User, error) {

	db, err := utils.OpenDatabase()

	if err != nil {
		return User{}, err
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM users WHERE username=? OR email=?")

	if err != nil {
		return User{}, err
	}

	rows, err := stmt.Query(login, login)

	if err != nil {
		return User{}, err
	}

	var user User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			return User{}, err
		}
	}

	return user, nil
}
