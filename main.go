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

func Insert(db *gorm.DB) {
	fmt.Println("Enter the id:")
	fmt.Scan(&id)
	fmt.Println("Enter the name:")
	fmt.Scan(&name)
	fmt.Println("Enter the domain:")
	fmt.Scan(&domain)
	data := Student{Id: id, Name: name, Domain: domain}
	db.Create(&data)
	fmt.Println("Value Inserted!!")
}

func Read(db *gorm.DB) {
	var student []Student //Slice of Student struct
	db.Find(&student)
	fmt.Println("id name domain")
	for _, student := range student {
		fmt.Printf("%d - %s - %s \n", student.Id, student.Name, student.Domain)
	}
}

func Update(db *gorm.DB) {
	fmt.Println("Enter the id:")
	fmt.Scan(&id)
	fmt.Println("Enter the name:")
	fmt.Scan(&name)
	fmt.Println("Enter new domain:")
	fmt.Scan(&domain)

	var student Student
	db.First(&student, id)
	student.Name = name
	student.Domain = domain
	db.Save(&student)
	fmt.Println("Updated!!!")
}
