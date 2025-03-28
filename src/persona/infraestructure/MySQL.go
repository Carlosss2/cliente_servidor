package infraestructure

import (
	"database/sql"
	"fmt"
	"long_short/src/persona/domain"
)

type MySQL struct {
	DB *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{DB: db}
}

func (mysql *MySQL) Save(persona domain.Persona) error {
	_, err := mysql.DB.Exec(
		"INSERT INTO persona (edad, nombre, sexo) VALUES (?, ?, ?)", 
		persona.Edad, persona.Nombre, persona.Sexo,
	)
	if err != nil {
		return fmt.Errorf("[MySQL] Error al guardar la persona: %w", err)
	}
	return nil
}
