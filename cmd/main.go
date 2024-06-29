package main

import (
	// "fmt"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/neem"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/api"

	"github.com/joho/godotenv"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		neem.Critial(err, "Error loading .env file")
	}
	
	api.Server()
}
