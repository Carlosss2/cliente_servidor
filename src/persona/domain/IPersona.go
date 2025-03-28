package domain

type IPersona interface{
	
	AddPerson(persona Persona) error
	GetnewPersonIsAdded() (bool, error)
	CountGender(bool)(int,error)
}