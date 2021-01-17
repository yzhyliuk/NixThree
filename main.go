package main

import "NixTwo/application"

func main() {
	App := application.InitApp("https://jsonplaceholder.typicode.com")
	App.Start()
}
