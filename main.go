package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// stdudent records
var id int
var name, domain string

type Student struct {
	Id     int
	Name   string
	Domain string
}

func main() {

}
// connecting to db with gorm pkg
func connectPostgresDB() *gorm.DB {
	connectTo := "host=localhost user=postgres password=******** dbname=databasename port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connectTo), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
