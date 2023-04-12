package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdavid87/sistema-reserva-turnos/internal/domain"
	"github.com/rdavid87/sistema-reserva-turnos/internal/turno"
	"github.com/rdavid87/sistema-reserva-turnos/pkg/util"
	"github.com/rdavid87/sistema-reserva-turnos/pkg/web"
)

type TurnoHandler interface {
	Add(c *gin.Context)
	Update(c *gin.Context)
	Patch(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	GetByDNI(c *gin.Context)
}

type turnoHandler struct {
	service turno.Service
}

func NewTurnoHandler(service turno.Service) TurnoHandler {
	return &turnoHandler{service}
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
	var turno domain.Turno
	if err := c.ShouldBindJSON(&turno); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	if err := h.service.Update(&turno); err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, turno)
}

func (h *turnoHandler) Patch(c *gin.Context) {
	var turno domain.Turno
	if err := c.ShouldBindJSON(&turno); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
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
	if turno.Id != id {
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
