package main

import (
	"fmt"
	"log"

	"github.com/syarifme/simpleshop/db"
)

func main() {
	_, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello")
}
