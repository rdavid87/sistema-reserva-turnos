package domain

type Paciente struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre" binding:"required"`
	Apellido  string `json:"apellido" binding:"required"`
	Domicilio string `json:"domicilio" binding:"required"`
	DNI       string `json:"dni" binding:"required"`
	FechaAlta string `json:"fecha_alta"`
}

type PacienteAbstract struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Domicilio string `json:"domicilio"`
	DNI       string `json:"dni"`
	FechaAlta string `json:"fecha_alta"`
}
