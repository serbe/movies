package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type config struct {
	Base struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
		Sslmode  bool   `json:"sslmode"`
	} `json:"postgresql"`
	Web struct {
		Port string `json:"port"`
	} `json:"web"`
}

func (app *application) getConfig() {
	c := config{}
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		panic(err)
	}
	app.config = c
	log.Println(app.config)
}
