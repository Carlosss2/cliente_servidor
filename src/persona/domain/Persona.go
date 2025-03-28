package domain

type Persona struct{
	Edad int32
	Nombre string
	Sexo bool


}

func NewPersona(Edad int32, Nombre string,Sexo bool)*Persona{
	return &Persona{Edad: Edad,Nombre: Nombre,Sexo: Sexo}
}