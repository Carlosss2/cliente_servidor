package controllers

import (
	"long_short/src/persona/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetNewPersonIsAddedController struct {
	useCase *application.GetNewPersonIsAddedUc
}

func NewGetNewPersonIsAddedController(useCase *application.GetNewPersonIsAddedUc)*GetNewPersonIsAddedController{
	return &GetNewPersonIsAddedController{useCase: useCase}

}

func (controller *GetNewPersonIsAddedController)Execute(ctx *gin.Context){
	newPersonAdded, err := controller.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"new_person_added": newPersonAdded})
}