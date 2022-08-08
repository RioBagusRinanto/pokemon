package main

import (
	"farmacare/pokemon/config"
	"farmacare/pokemon/routes"
)

func main() {
	config.InitDB()

	e := routes.New()

	port := "4100"

	e.Logger.Fatal(e.Start(":" + port))

}
