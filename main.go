package main

import (
	"NixTwo/application"
	"fmt"
)

//import "NixTwo/application"

func main() {
	app, err := application.NewApp()
	if err != nil {
		fmt.Printf(err.Error())
	}
	app.Start()
}
