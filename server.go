package main

import (
	"fmt"
	"log"
	connection "test-mssql/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	connection.Connection()
	query := "exec [getStatusOrderIn]"

	rows, err := connection.Conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var status string
		var count int32

		err := rows.Scan(&status, &count)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("found %s : %d\n", status, count)
		}
	}
	// query := "exec [getStatusOrderIn] @warehouse='U'"

	// rows, err := conn.Query(query)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	// var id string
	// 	var status string
	// 	var count int32

	// 	err := rows.Scan(&status, &count)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	} else {
	// 		log.Printf("Found %s %d\n", status, count)
	// 	}
	// }

	// app := fiber.New()

	// app.Get("/", helloWorld)
	// app.Listen(":8080")
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
