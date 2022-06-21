package main

import (
	"fmt"

	route2 "github.com/codeedu/imersaofsfc2-simulator/application/route"
)

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPostions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Println(stringJson[1])
}
