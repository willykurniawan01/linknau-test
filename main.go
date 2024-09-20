package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/willykurniawan01/linknau-test/app/router"
	"github.com/willykurniawan01/linknau-test/app/services"
)

func main() {
	log.Println("== ⁠Structs: Define a struct in Go to represent a Person with Name and Age fields. ==")
	person := services.Person{}
	person.Name = "Willy kurniawan"
	person.Age = 24
	log.Printf("My Name is %s ,iam %d years old", person.Name, person.Age)

	log.Println("== Interfaces : Explain what an interface is in Go and provide an example of a simple interface. ==")
	car := services.Car{Model: "Toyota"}
	bike := services.Bike{Brand: "Yamaha"}
	services.OperateVehicle(car)
	services.OperateVehicle(bike)

	log.Println("== ⁠Package Management: Explain how Go's package management system works and how you would import a third-party package in a Go project. ==")
	services.FetchDataFromAPI()

	log.Println("== ⁠Authentication: Implement a secure authentication and authorization mechanism using JWT authentication ==")
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	router.Route(r)
	r.Run(":8080")
}
