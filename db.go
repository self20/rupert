package rupert

import (
	"crypto/sha256"
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"math/rand"
	"time"
)

var (
	db *sql.DB
)

type User struct {
	UserID  int64  `db:"user_id"`
	Name    string `db:"name"`
	Created int64  `db:"created_at"`
	Salt    string `db:"salt"`
	Hash    []byte `db:"key"`
}

func randString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func newUser(name, password string) User {
	salt := randString(20)
	return User{
		Created: time.Now().UnixNano(),
		Name:    name,
		Salt:    salt,
		Hash:    computeHash(password, salt),
	}
}

func computeHash(password, salt string) []byte {
	h := sha256.New()
	io.WriteString(h, password)
	io.WriteString(h, salt)
	hash := h.Sum(nil)
	return hash
}

func initDb() *sql.DB {
	db, err := sql.Open("sqlite3", "/tmp/rupert_db.sqlite")
	checkErr(err, "sql.Open Failed to open database")
	return db
}
