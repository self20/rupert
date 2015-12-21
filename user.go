package rupert

import (
	"log"
	"time"
)

type (
	User struct {
		UserID    int
		Username  string
		Hash      []byte
		Salt      string
		Enabled   bool
		CreatedOn time.Time
		UpdatedOn time.Time
	}
)

var (
	queryCreateUser = `
		INSERT INTO users (username, hash, salt, enabled)
		VALUES (:username, :hash, :salt, :enabled)
	`
	queryUserByID = `
		SELECT user_id, username, hash, salt, enabled, created_on, updated_on
		FROM users
		WHERE user_id=$1
	`

	queryUserByName = `
		SELECT user_id, username, hash, salt, enabled, created_on, updated_on
		FROM users
		WHERE username=$1
	`
)

func NewUser(name, password string) User {
	salt := randString(20)
	return User{
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
		Username:  name,
		Salt:      salt,
		Hash:      computeHash(password, salt),
	}
}

func GetUserByID(user_id int) (*User, error) {
	user := User{}
	err := db.Get(&user, queryUserByID, user_id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByName(user_name string) (*User, error) {
	user := User{}
	err := db.Get(&user, queryUserByName, user_name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(username string, password string, enabled bool) (*User, error) {
	log.Println("1")
	user := NewUser(username, password)
	user.Enabled = enabled
	log.Println("2")
	tx := db.MustBegin()
	log.Println("3")
	_, err := tx.NamedExec(queryCreateUser, &user)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return GetUserByName(username)
}
