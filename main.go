package main

import (
	_ "encoding/json"
	"fmt"
	"github.com/umeshdhaked/go_rest/apis/apitest"
	"github.com/umeshdhaked/go_rest/apis/users"
	"log"
	_ "math/rand"
	"net/http"
	_ "strconv"
)

func main() {
	port := 8080
	log.Printf("Starting server on port:%v\n", port)

	handlers()

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}

func handlers() {
	http.HandleFunc("/", apitest.TestApi) // ANY METHOD
	http.HandleFunc("/user/createUser",users.CreateUser)  //POST METHOD ONLY
}

