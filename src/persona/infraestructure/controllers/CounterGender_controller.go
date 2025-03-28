package controllers

import (
	"context"
	"net/http"
	"time"

	"long_short/src/persona/application"
	"github.com/gin-gonic/gin"
)

type CountGenderController struct {
	useCase *application.CountGenderUc
}

func NewCountGenderUc(useCase *application.CountGenderUc) *CountGenderController {
	return &CountGenderController{useCase: useCase}
}

func (controller *CountGenderController) Execute(c *gin.Context) {
	// Obtener el parámetro de la consulta (sexo)
	sexo := c.Param("sexo")

	var sexoBool bool
	if sexo == "true" {
		sexoBool = true
	} else if sexo == "false" {
		sexoBool = false
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor de sexo inválido. Debe ser 'true' o 'false'."})
		return
	}

	// Establecer el encabezado de respuesta
	c.Header("Content-Type", "application/json")

	// Crear un contexto con timeout de 10 segundos
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	// Obtener el conteo inicial
	initialCount, err := controller.useCase.Execute(sexoBool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Crear un ticker para verificar cambios cada 1 segundo
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// Si el tiempo expira, devolver el conteo sin cambios
			c.JSON(http.StatusOK, gin.H{"conteo": initialCount, "timeout": true})
			return
		case <-ticker.C:
			// Verificar si hay cambios en el conteo
			currentCount, err := controller.useCase.Execute(sexoBool)
			if err == nil && currentCount != initialCount {
				// Si el conteo ha cambiado, responder de inmediato
				c.JSON(http.StatusOK, gin.H{"conteo": currentCount, "updated": true})
				c.Writer.Flush()
				return
			}
		}
	}
}
