package domain

type IPersona interface{
	
	AddPerson(persona Persona) error
	GetAll()([]Persona,error) 
}