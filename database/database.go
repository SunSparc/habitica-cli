package database

import (
	"io/ioutil"
	"log"
)

type Database struct {
	//data []something
	filename string
}

func InitializeDatabase(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		if err.Error() == "no such file or directory" {
			log.Println()
		}
		log.Println("ioutil.ReadFile:", err)
	}
	if len(bytes) <= 0 {
		log.Println("the database file is empty")
	}
}

func NewDatabase(filename string) *Database {
	return &Database{}
}

// json.marshalindent
