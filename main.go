package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello Wrld")
	myName := "Satyam"
	fmt.Println(myName)

	app := fiber.New()

	log.Fatal(app.Listen(":4000"))

}