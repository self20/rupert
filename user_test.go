package rupert

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
	"time"
)

func checkTestErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err.Error())
	}

}
func TestCreateUser(t *testing.T) {
	user, err := CreateUser("test", "test", true)
	checkTestErr(t, err)
	if user.Enabled != true || user.Username != "test" {
		t.Error("Invalid value", err)
	}
	if user.UserID <= 0 {
		t.Error("Invalid userid", err)
	}
}

func init() {
	if config.Testing {
		db_name := fmt.Sprintf("rupert_test_%d", time.Now().Unix())
		//		log.Println("Dropping DB")
		//		err := exec.Command("dropdb", db_name).Start()
		//		checkErr(err, "Failed to droptest database")
		//		time.Sleep(2 * time.Second)
		log.Println("Creating DB")
		err := exec.Command("createdb", db_name).Start()
		checkErr(err, "Failed to create test database")
		time.Sleep(2 * time.Second)
		db = initDb(fmt.Sprintf("dbname=%s sslmode=disable", db_name))
		db.MustExec(Schema)
		//defer db.Close()
		//checkErr(db.Ping(), "Failed to connect to test database as configured adress")
		//forum.Initialize()
	}
}
