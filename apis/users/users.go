package users

import (
	"encoding/json"
	"fmt"
	_ "math/rand"
	"net/http"
	_ "strconv"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var data Data
		err := json.NewDecoder(req.Body).Decode(&data)
		if err == nil {
			res.Header().Set("Content-Type","application/json")
			json.NewEncoder(res).Encode(data)

		} else {
			fmt.Fprintf(res, "Internal Server Error!! Unable to decode data")
		}
	} else {
		http.Error(res, "Not Found!!", http.StatusNotFound)
	}
}

type Data struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

type CustomError struct {
	message string `json:"message"`
}

func (d Data) String() string {
	return fmt.Sprintf("{ key1: %v, key2: %v }", d.Key1, d.Key2)
}
