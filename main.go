package main

import (
	"github.com/ariwanss/CvBackendGo/repository"
	"github.com/ariwanss/CvBackendGo/router"
	"github.com/ariwanss/CvBackendGo/service"
	"github.com/ariwanss/CvBackendGo/setup"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	loadEnv()
	repository.ConnectDb("Sandbox")
	service.StartService()
	router.SetupRouter()
	setup.RunSetup()
	router.Router.Run("localhost:5000")
}
