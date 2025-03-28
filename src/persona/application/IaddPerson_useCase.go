package application

import (
	"long_short/src/persona/domain"
)

type IaddPerson struct {
	db domain.IPersona // Se espera que IPersona tenga el método AddPerson
}

func NewIaddPerson(db domain.IPersona) *IaddPerson {
	return &IaddPerson{db: db}
}

// Ahora la función usa correctamente el parámetro persona
func (uc *IaddPerson) Execute(persona domain.Persona) error {
	return uc.db.AddPerson(persona) // Usa el objeto persona recibido
}
