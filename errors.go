package rupert

import (
	log "github.com/Sirupsen/logrus"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
