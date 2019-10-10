package main

import (
	"auth/models"
	"auth/routes"
)

func init()  {
	models.Setup()
}

func main()  {
	r := routes.Setup()
	r.Run(":3000")
}