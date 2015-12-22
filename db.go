package rupert

import (
	"crypto/sha256"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io"
	"math/rand"
)

var (
	db *sqlx.DB
)

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

func computeHash(password, salt string) []byte {
	h := sha256.New()
	io.WriteString(h, password)
	io.WriteString(h, salt)
	hash := h.Sum(nil)
	return hash
}

func initDb(db_dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", db_dsn)
	checkErr(err, "sql.Open Failed to open database")
	return db
}
