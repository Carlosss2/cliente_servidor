package controllers



import (
	"net/http"
	
	"long_short/src/persona/application"
	"github.com/gin-gonic/gin"
)
type CountGenderController struct {
	useCase *application.CountGenderUc
}

func NewCountGenderUc(useCase *application.CountGenderUc)*CountGenderController{
	return &CountGenderController{useCase: useCase}
   }
   
   func (controller *CountGenderController) Execute(c *gin.Context){
	   // Obtener el parámetro de la consulta (sexo) de la solicitud
	   sexo := c.Param("sexo") // "true" es el valor por defecto si no se pasa nada
   
	   var sexoBool bool
	   if sexo == "true" {
		   sexoBool = true
	   } else if sexo == "false" {
		   sexoBool = false
	   } else {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "Valor de sexo inválido. Debe ser 'true' o 'false'."})
		   return
	   }
   
	   // Llamar al caso de uso y pasar el parámetro sexo
	   count, err := controller.useCase.Execute(sexoBool)
	   if err != nil {
		   c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		   return
	   }
   
	   // Devolver el conteo
	   c.JSON(http.StatusOK, gin.H{"conteo": count})
   }