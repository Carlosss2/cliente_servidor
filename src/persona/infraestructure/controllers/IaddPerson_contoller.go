package controllers

import (
	"long_short/src/persona/application"
	"long_short/src/persona/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IaddPersonController struct {
	useCaseAddPerson *application.IaddPerson
}

func NewIaddPersonController(useCaseAddPerson *application.IaddPerson) *IaddPersonController {
	return &IaddPersonController{useCaseAddPerson: useCaseAddPerson}
}

func (createPersona *IaddPersonController) Create(c *gin.Context) {
	var persona domain.Persona

	if err := c.ShouldBindJSON(&persona); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createPersona.useCaseAddPerson.Execute(persona) // Pasar `persona` correctamente
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Persona registrada"})
}
