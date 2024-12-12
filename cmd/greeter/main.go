package main

import (
	app "clean_arch_basic_example/internal"
)

// run app
func main() {

	// init environments
	app.InitEnv(".env")

	// run app
	app.Run()
}
