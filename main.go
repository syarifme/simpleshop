package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syarifme/simpleshop/controllers"
	"github.com/syarifme/simpleshop/db"
	userRepo "github.com/syarifme/simpleshop/repositories/user"
)

func main() {
	r := mux.NewRouter()
	server := new(http.Server)

	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	response := userRepo.UserRepo(conn).GetAll()
	fmt.Printf("%+v", response)

	r.HandleFunc("/login/{username}", controllers.HandleLogin)

	server.Addr = ":9000"
	server.Handler = r
	log.Println("Starting new server at :9000")
	server.ListenAndServe()
}
