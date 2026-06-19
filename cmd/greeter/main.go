package main

import (
	app "github.com/iqhater/clean-arch-basic-example/internal"
)

// run app
func main() {

	// init environments
	app.InitEnv(".env")

	// run app
	app.Run()
}
