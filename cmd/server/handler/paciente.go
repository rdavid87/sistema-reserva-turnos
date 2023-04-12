package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdavid87/sistema-reserva-turnos/internal/domain"
	"github.com/rdavid87/sistema-reserva-turnos/internal/paciente"
	"github.com/rdavid87/sistema-reserva-turnos/pkg/util"
	"github.com/rdavid87/sistema-reserva-turnos/pkg/web"
)

type PacienteHandler interface {
	Add(c *gin.Context)
	Update(c *gin.Context)
	Patch(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
}

type pacienteHandler struct {
	service paciente.Service
}

func NewPacienteHandler(service paciente.Service) PacienteHandler {
	return &pacienteHandler{service}
}

func (h *pacienteHandler) Add(c *gin.Context) {
	var paciente domain.Paciente
	if err := c.ShouldBindJSON(&paciente); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	id, err := h.service.Add(&paciente)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusCreated, id)
}

func (h *pacienteHandler) Update(c *gin.Context) {
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
	if existe == nil {
		web.NotFound(c)
		return
	}

	var paciente domain.Paciente
	if err := c.ShouldBindJSON(&paciente); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	paciente.Id = existe.Id

	if err := h.service.Update(&paciente); err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, paciente)
}

func (h *pacienteHandler) Patch(c *gin.Context) {
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

	var paciente domain.PacienteAbstract
	if err := c.ShouldBindJSON(&paciente); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}

	if paciente.Apellido != "" {
		existe.Apellido = paciente.Apellido
	}

	if paciente.Nombre != "" {
		existe.Nombre = paciente.Nombre
	}

	if paciente.DNI != "" {
		existe.DNI = paciente.DNI
	}

	if paciente.Domicilio != "" {
		existe.Domicilio = paciente.Domicilio
	}

	if err := h.service.Update(existe); err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, existe)
}

func (h *pacienteHandler) Delete(c *gin.Context) {
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

func (h *pacienteHandler) GetByID(c *gin.Context) {
	id, err := util.GetIdFromParam(c)
	if err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	paciente, err := h.service.GetByID(id)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if paciente.Id != id {
		web.NotFound(c)
		return
	}
	web.Success(c, http.StatusOK, paciente)
}

func (h *pacienteHandler) GetAll(c *gin.Context) {
	dentists, err := h.service.GetAll()
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, dentists)
}
