package application

import "long_short/src/persona/domain"

type GetNewPersonIsAddedUc struct {
	db domain.IPersona
}

func NewGetNewPersonIsAddedUc(db domain.IPersona) *GetNewPersonIsAddedUc {
	return &GetNewPersonIsAddedUc{db: db}
}

func (useCase *GetNewPersonIsAddedUc) Execute() (bool, error) {
	return useCase.db.GetnewPersonIsAdded()
}