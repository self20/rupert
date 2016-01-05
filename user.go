package rupert

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type (
	User struct {
		UserID    int       `db:"user_id"`
		Username  string    `db:"username"`
		Hash      []byte    `db:"hash"`
		Salt      string    `db:"salt"`
		Enabled   bool      `db:"enabled"`
		CreatedOn time.Time `db:"created_on"`
		UpdatedOn time.Time `db:"updated_on"`
	}
)

var (
	queryCreateUser = `
		INSERT INTO users
			(username, hash, salt, enabled)
		VALUES
			(:username, :hash, :salt, :enabled)
	`
	queryUserByID = `
		SELECT
			user_id, username, hash, salt, enabled, created_on, updated_on
		FROM
			users
		WHERE
			user_id=$1
	`

	queryUserByName = `
		SELECT
			user_id, username, hash, salt, enabled, created_on, updated_on
		FROM
			users
		WHERE
			username=$1
	`
	queryUserUpdate = `
		UPDATE
			users
		SET
			username = :username,
			hash = :hash,
			salt = :salt,
			enabled = :enabled,
			updated_on = :updated_on
		WHERE
			user_id = :user_id
	`
)

func UserNew(name, password string) User {
	salt := randString(20)
	return User{
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
		Username:  name,
		Salt:      salt,
		Hash:      computeHash(password, salt),
		Enabled:   false,
	}
}

func UserSave(db *sqlx.DB, user *User) error {
	user.UpdatedOn = time.Now()
	_, err := db.NamedExec(queryUserUpdate, user)
	return err
}

func UserGetByID(db *sqlx.DB, user_id int) (*User, error) {
	var user User
	err := db.Get(&user, queryUserByID, user_id)
	return &user, err
}

func UserGetByName(db *sqlx.DB, user_name string) (*User, error) {
	user := User{}
	err := db.Get(&user, queryUserByName, user_name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserDelete(db *sqlx.DB, user_id int) error {
	_, err := db.Exec("DELETE FROM users WHERE user_id = $1", user_id)
	return err
}

func UserCreate(db *sqlx.DB, username string, password string, enabled bool) (*User, error) {
	user := UserNew(username, password)
	user.Enabled = enabled
	tx := db.MustBegin()
	_, err := tx.NamedExec(queryCreateUser, &user)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return UserGetByName(db, username)
}
