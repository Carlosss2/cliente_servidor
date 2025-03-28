package infraestructure

import (
	"database/sql"
	"fmt"
	"long_short/src/persona/domain"
)

type MySQL struct {
	DB *sql.DB
	lastCount int
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{DB: db}
}

func (mysql *MySQL) AddPerson(persona domain.Persona) error {
	_, err := mysql.DB.Exec(
		"INSERT INTO persona (edad, nombre, sexo) VALUES (?, ?, ?)", 
		persona.Edad, persona.Nombre, persona.Sexo,
	)
	if err != nil {
		return fmt.Errorf("[MySQL] Error al guardar la persona: %w", err)
	}
	return nil
}

func (sql *MySQL)GetnewPersonIsAdded() (bool, error){

	var count int
	err := sql.DB.QueryRow("SELECT COUNT(*) FROM personas").Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error obteniendo el conteo de personas: %v", err)
	}

	if count > sql.lastCount {
		// Si el conteo actual es mayor que el anterior, significa que se ha agregado una persona
		// Actualizamos lastCount para mantener el conteo más reciente
		sql.lastCount = count
		return true, nil
	}

	return false, nil
}