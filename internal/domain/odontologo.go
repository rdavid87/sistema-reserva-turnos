package domain

type Odontologo struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre" binding:"required"`
	Apellido  string `json:"apellido" binding:"required"`
	Matricula string `json:"matricula" binding:"required"`
}

type OdontologoAbstract struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Matricula string `json:"matricula"`
}
