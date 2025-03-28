package dependencies

import (
	"database/sql"
	"fmt"

	"long_short/src/core"
	"long_short/src/persona/application"
	"long_short/src/persona/infraestructure"
	"long_short/src/persona/infraestructure/controllers"
)

var (
	mySQL infraestructure.MySQL
	db    *sql.DB
)

func Init() {
	var err error
	db, err = core.ConnectToDB() // Corrige la asignación de `db`

	if err != nil {
		fmt.Println("server error")
		return
	}

	mySQL = *infraestructure.NewMySQL(db)
}

func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Conexión a la base de datos cerrada.")
	}
}

// Corrige el nombre de la función
func GetCreatePersonController() *controllers.IaddPersonController {
	caseCreatePerson := application.NewIaddPerson(&mySQL) // Asegurar que NewIaddPerson devuelva un puntero a la estructura correcta
	return controllers.NewIaddPersonController(caseCreatePerson)
}
func GetNewPersonIsAddedController()*controllers.GetNewPersonIsAddedController{
	useCase := application.NewGetNewPersonIsAddedUc(&mySQL)
	return controllers.NewGetNewPersonIsAddedController(useCase)
}