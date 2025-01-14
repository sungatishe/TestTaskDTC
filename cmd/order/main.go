package main

import (
	_ "TestTask/docs"
	"TestTask/internal/app"
)

// @SecurityDefinitions.apiKey ApiKeyAuth
// @Description API Key authorization
// @In header
// @Name Authorization

func main() {
	app.Run()
}
