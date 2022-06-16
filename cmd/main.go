package main

import (
	"comento_git_practice/app"
	"log"
)

func main() {
	log.Fatal(app.NewEcho(*app.DefaultConfig()).Start(":8080"))
}
