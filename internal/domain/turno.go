package domain

type Turno struct {
	Id          int        `json:"id"`
	Odontologo  Odontologo `json:"odontologo_id" binding:"required"`
	Paciente    Paciente   `json:"paciente_id" binding:"required"`
	Fecha       string     `json:"fecha" binding:"required"`
	Hora        string     `json:"hora" binding:"required"`
	Descripcion string     `json:"descripcion"`
}

type TurnoAbstract struct {
	Id           int    `json:"id"`
	OdontologoId int    `json:"odontologo_id"`
	PacienteId   int    `json:"paciente_id"`
	Fecha        string `json:"fecha"`
	Hora         string `json:"hora"`
	Descripcion  string `json:"descripcion"`
}

type TurnoResponse struct {
	Id          int
	Paciente    string
	Odontologo  string
	Fecha       string
	Hora        string
	Descripcion string
}
