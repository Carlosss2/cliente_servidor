package routes

import (
	"long_short/src/persona/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/persons")
	getNewPersinIsAddedController := dependencies.GetNewPersonIsAddedController().Execute
	iAddPerson := dependencies.GetCreatePersonController().Create
	getCountGenderController := dependencies.GetCountGenderController().Execute
	routes.POST("/",iAddPerson)
	routes.GET("/InewPersonIsAdded",getNewPersinIsAddedController)
	routes.GET("/CountGender/:sexo",getCountGenderController)
}