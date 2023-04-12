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
	OdontologoId int    `json:"odontologo_id" binding:"required"`
	PacienteId   int    `json:"paciente_id" binding:"required"`
	Fecha        string `json:"fecha" binding:"required"`
	Hora         string `json:"hora" binding:"required"`
	Descripcion  string `json:"descripcion"`
}

type TurnoDetalle struct {
	Id          int    `json:"id"`
	Paciente    string `json:"paciente" binding:"required"`
	Odontologo  string `json:"odontologo" binding:"required"`
	Fecha       string `json:"fecha" binding:"required"`
	Hora        string `json:"hora" binding:"required"`
	Descripcion string `json:"descripcion"`
}
