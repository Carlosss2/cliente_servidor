package routes

import (
	"long_short/src/persona/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/persons")

	iAddPerson := dependencies.GetCreatePersonController().Create

	routes.POST("/",iAddPerson)
}