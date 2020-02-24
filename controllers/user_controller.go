package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/syarifme/simpleshop/models"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", user)

	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(fmt.Sprintf("Hello %v", params["username"])))
}
