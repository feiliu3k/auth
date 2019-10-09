package main

import (
	"auth/models"
	"auth/router"
)

func init()  {
	models.Setup()
}

func main()  {
	r := router.Setup()
	r.Run(":3000")
}