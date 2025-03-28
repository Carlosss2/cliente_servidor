package controllers

import "long_short/src/persona/application"

type CountGenderUc struct {
	useCase *application.CountGenderUc
}

func NewCountGenderUc(useCase *application.CountGenderUc) *CountGenderUc {
	return &CountGenderUc{useCase: useCase}
}

