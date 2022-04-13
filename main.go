package main

import (
	"github.com/Feggah/banking-api/app"
	"github.com/Feggah/banking-api/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
