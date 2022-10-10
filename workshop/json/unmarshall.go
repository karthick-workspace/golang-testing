package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Instagram string `json:"instagram"`
	Facebook  string `json:"facebook"`
}

func main() {
	jsonFile, err := os.Open("./workshop/json/users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	log.Println("File descriptor successfully created")

	var users Users

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteValue, &users)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(users)

	// Option 2: Unmarshall using map

	var result map[string]interface{}

	err = json.Unmarshal(byteValue, &result)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("===================================================")
	log.Println(result["users"])

	log.Println("===================================================")
	log.Println(result["users"])

}
