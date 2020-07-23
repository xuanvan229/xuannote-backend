package main

import (
	"note/app"
)

// comment
func main() {
	// config := config.GetPostGresConfig()
	app := &app.App{}
	app.Initialize()
	app.Router.Run(":1323")
}
