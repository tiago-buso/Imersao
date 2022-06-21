package main

import (
	"fmt"
	"log"

	route2 "github.com/codeedu/imersaofsfc2-simulator/application/route"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPostions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Println(stringJson[1])
}
