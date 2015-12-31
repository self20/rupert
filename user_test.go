package rupert

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"testing"
	"time"
)

var (
	test_user_name = "test_user"
)

func checkTestNil(t *testing.T, i interface{}) {
	if i == nil {
		t.Error("Nil value found")
	}
}

func TestUsers(t *testing.T) {
	db := NewTestDB()
	// Create user
	user, err := UserCreate(db, test_user_name, "test", true)
	if err != nil {
		t.Error(err.Error())
	}
	checkTestNil(t, user)
	if user.Enabled != true || user.Username != test_user_name {
		t.Error("Invalid value", err)
	}
	if user.UserID <= 0 {
		t.Error("Invalid userid", err)
	}

	// Fetch by name
	user_name, err := UserGetByName(db, test_user_name)
	if err != nil {
		t.Error("Failed to retrieve newly created user")
	}
	if user_name.UserID <= 0 {
		t.Error("Invalid user id returned")
	}

	// Fetch by name
	user_id, err := UserGetByID(db, user_name.UserID)
	if err != nil {
		t.Error("Failed to retrieve newly created user", err.Error())
	}
	if user_id.UserID <= 0 {
		t.Error("Invalid user id returned")
	}

	// Save user changes
	time.Sleep(1 * time.Second) // make sure the time stamp differs
	user_id.Username = test_user_name + test_user_name
	err = UserSave(db, user_id)
	if err != nil {
		t.Error("Failed to update user", err)
	}

	// Fetch by name
	user_id2, err := UserGetByID(db, user_name.UserID)
	if err != nil {
		t.Error("Failed to retrieve newly created user", err.Error())
	}
	if user_id2.Username != test_user_name+test_user_name {
		t.Error("Invalid user name returned")
	}

	if user_id2.UpdatedOn.Unix() <= user_id2.CreatedOn.Unix() {
		t.Error("Invalid update date returned")
	}

	// Delete user
	err = UserDelete(db, user_id.UserID)
	if err != nil {
		t.Error("Could not delete user", err.Error())
	}
	_, err = UserGetByID(db, user_name.UserID)
	if err == nil {
		t.Error("Could not fully delete user", err.Error())
	}
}

func NewTestDB() *sqlx.DB {
	// Allow a alternate DB to be specified for testing
	db_name := os.Getenv("RUPERT_TEST_DB")
	if db_name == "" {
		db_name = "rupert_test"
	}
	db := initDb(fmt.Sprintf("dbname=%s sslmode=disable", db_name))

	// Re-initialize a empty DB
	db.MustExec("drop schema public cascade;")
	db.MustExec("create schema public;")
	db.MustExec(Schema)
	return db
}
