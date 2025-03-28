package main

import (
	"long_short/src/persona/infraestructure/dependencies"
	"long_short/src/persona/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dependencies.Init()

	defer dependencies.CloseDB()

	r := gin.Default()

	routes.Routes(r)
	r.Run()
}