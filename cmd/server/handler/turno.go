package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdavid87/sistema-reserva-turnos/internal/domain"
	"github.com/rdavid87/sistema-reserva-turnos/internal/odontologo"
	"github.com/rdavid87/sistema-reserva-turnos/internal/paciente"
	"github.com/rdavid87/sistema-reserva-turnos/internal/turno"
	"github.com/rdavid87/sistema-reserva-turnos/pkg/util"
	"github.com/rdavid87/sistema-reserva-turnos/pkg/web"
)

type TurnoHandler interface {
	Add(c *gin.Context)
	AddByDniMatricula(c *gin.Context)
	Update(c *gin.Context)
	Patch(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	GetByDNI(c *gin.Context)
}

type turnoHandler struct {
	service           turno.Service
	serviceOdontologo odontologo.Service
	servicePaciente   paciente.Service
}

func NewTurnoHandler(service turno.Service, serviceOdontologo odontologo.Service, servicePaciente paciente.Service) TurnoHandler {
	return &turnoHandler{service, serviceOdontologo, servicePaciente}
}

func (h *turnoHandler) Add(c *gin.Context) {
	var turno domain.TurnoAbstract
	if err := c.ShouldBindJSON(&turno); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	id, err := h.service.Add(&turno)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusCreated, id)
}

func (h *turnoHandler) Update(c *gin.Context) {
	id, err := util.GetIdFromParam(c)
	if err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	existeTurno, err := h.service.GetByID(id)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if existeTurno == nil {
		web.NotFound(c)
		return
	}

	var turno domain.TurnoAbstract
	if err := c.ShouldBindJSON(&turno); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}

	turno.Id = existeTurno.Id

	if err := h.service.Update(&turno); err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, turno)
}

func (h *turnoHandler) Patch(c *gin.Context) {
	id, err := util.GetIdFromParam(c)
	if err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	existe, err := h.service.GetByID(id)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if existe.Id == 0 {
		web.NotFound(c)
		return
	}

	var turno domain.TurnoAbstract
	if err := c.ShouldBindJSON(&turno); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}

	turno.Id = existe.Id

	if turno.Descripcion == "" {
		turno.Descripcion = existe.Descripcion
	}

	if turno.Fecha == "" {
		turno.Fecha = existe.Fecha
	}

	if turno.Hora == "" {
		turno.Hora = existe.Hora
	}

	if turno.OdontologoId == 0 {
		turno.OdontologoId = existe.Odontologo.Id
	}

	if turno.PacienteId == 0 {
		turno.OdontologoId = existe.Paciente.Id
	}

	if err := h.service.Update(&turno); err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, turno)
}

func (h *turnoHandler) Delete(c *gin.Context) {
	id, err := util.GetIdFromParam(c)
	if err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	if err := h.service.Delete(id); err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, nil)
}

func (h *turnoHandler) GetByID(c *gin.Context) {
	id, err := util.GetIdFromParam(c)
	if err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	turno, err := h.service.GetByID(id)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if turno == nil {
		web.NotFound(c)
		return
	}
	web.Success(c, http.StatusOK, turno)
}

func (h *turnoHandler) GetAll(c *gin.Context) {
	dentists, err := h.service.GetAll()
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, dentists)
}

func (h *turnoHandler) GetByDNI(c *gin.Context) {
	dni := util.GetDNIFromParam(c)
	if dni == "" {
		web.Failure(c, http.StatusBadRequest, errors.New("DNI faltante"))
		return
	}
	turno, err := h.service.GetByDNI(dni)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if turno.Id == 0 {
		web.NotFound(c)
		return
	}
	web.Success(c, http.StatusOK, turno)
}

func (h *turnoHandler) AddByDniMatricula(c *gin.Context) {
	dni, matricula := util.GetDniMatriculaFromParam(c)
	if dni == "" || matricula == "" {
		web.Failure(c, http.StatusBadRequest, errors.New("parametro Dni o Matricula no encontrado"))
		return
	}

	var turno domain.TurnoAbstract
	if err := c.ShouldBindJSON(&turno); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}

	if turno.Fecha == "" {
		web.Failure(c, http.StatusBadRequest, errors.New("fecha requerida"))
		return
	}

	if turno.Hora == "" {
		web.Failure(c, http.StatusBadRequest, errors.New("hora requerida"))
		return
	}

	odontologo, err := h.serviceOdontologo.GetByMatricula(matricula)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}

	if odontologo.Matricula != matricula {
		web.NotFound(c)
		return
	}

	paciente, err := h.servicePaciente.GetByDNI(dni)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if paciente.DNI != dni {
		web.NotFound(c)
		return
	}

	turno.OdontologoId = odontologo.Id
	turno.PacienteId = paciente.Id

	id, err := h.service.Add(&turno)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusCreated, id)
}
