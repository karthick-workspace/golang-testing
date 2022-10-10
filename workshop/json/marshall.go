package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Professor struct {
	Name       string     `json:"name"`
	ScienceID  int        `json:"science_id"`
	IsWorking  bool       `json:"is_working"`
	University University `json:"university"`
}

type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	prof1 := Professor{
		Name:      "Bob",
		ScienceID: 12312,
		IsWorking: false,
		University: University{
			Name: "KPM",
			City: "KPM",
		},
	}

	byteArr, err := json.Marshal(prof1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byteArr))
	err = os.WriteFile("./workshop/json/output.json", byteArr, 0664)
	if err != nil {
		log.Fatal(err)
	}

}
