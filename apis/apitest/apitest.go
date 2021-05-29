package apitest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func TestApi(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		log.Printf("serving '/'")

		res.Write([]byte("It's Working fine !!\n"))
		fmt.Fprintln(res,"Howdy :) ")

	} else {

		// 404 Not found
		res.Header().Set("Content-Type","application/json")
		res.WriteHeader(http.StatusNotFound)

		json.NewEncoder(res).Encode(struct {
			Message string `json:"message"`
		}{Message: "Only GET request accepted"})
	}
}