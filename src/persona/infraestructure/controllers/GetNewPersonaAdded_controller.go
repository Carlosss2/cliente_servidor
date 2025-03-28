package controllers

import (
	"context"
	"long_short/src/persona/application"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GetNewPersonIsAddedController struct {
	useCase *application.GetNewPersonIsAddedUc
}

func NewGetNewPersonIsAddedController(useCase *application.GetNewPersonIsAddedUc) *GetNewPersonIsAddedController {
	return &GetNewPersonIsAddedController{useCase: useCase}
}

func (controller *GetNewPersonIsAddedController) Execute(c *gin.Context) {
	// Obtener el parámetro ID de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe de ser un número entero"})
		return
	}

	// Establecer el encabezado de respuesta
	c.Header("Content-Type", "application/json")

	// Crear un contexto con timeout de 10 segundos
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	// Crear un ticker para verificar cada segundo
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// Si el tiempo expira, devolver un timeout
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timed out"})
			return
		case <-ticker.C:
			// Verificar si hay una nueva persona añadida
			newPersonAdded, err := controller.useCase.Execute()
			if err == nil && newPersonAdded {
				c.JSON(http.StatusOK, gin.H{
					"new_person_added": newPersonAdded,
					"id":               id,
				})
				c.Writer.Flush()
				return
			}
		}
	}
}
