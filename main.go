package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/routes"
)

func main() {
	loadData()
	r := routes.GetRoutes()
	log.Fatal(r.Run())
}

func loadData() {
	f, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = json.Unmarshal(f, &models.MyResume)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
