package application

import "long_short/src/persona/domain"

type CountGenderUc struct {
	db domain.IPersona
}

func NewCountGenderUc(db domain.IPersona) *CountGenderUc {
	return &CountGenderUc{db: db}
}

