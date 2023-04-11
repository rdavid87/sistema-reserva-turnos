package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdavid87/sistema-reserva-turnos/internal/domain"
	"github.com/rdavid87/sistema-reserva-turnos/internal/odontologo"
	"github.com/rdavid87/sistema-reserva-turnos/pkg/util"
	"github.com/rdavid87/sistema-reserva-turnos/pkg/web"
)

type OdontologoHandler interface {
	Add(c *gin.Context)
	Update(c *gin.Context)
	Patch(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
}

type odontologoHandler struct {
	service odontologo.Service
}

func NewOdontologoHandler(service odontologo.Service) OdontologoHandler {
	return &odontologoHandler{service}
}

func (h *odontologoHandler) Add(c *gin.Context) {
	var odontologo domain.Odontologo
	if err := c.ShouldBindJSON(&odontologo); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	id, err := h.service.Add(&odontologo)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusCreated, id)
}

func (h *odontologoHandler) Update(c *gin.Context) {
	var odontologo domain.Odontologo
	if err := c.ShouldBindJSON(&odontologo); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	if err := h.service.Update(&odontologo); err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, odontologo)
}

func (h *odontologoHandler) Patch(c *gin.Context) {
	var odontologo domain.Odontologo
	if err := c.ShouldBindJSON(&odontologo); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	if err := h.service.Update(&odontologo); err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, odontologo)
}

func (h *odontologoHandler) Delete(c *gin.Context) {
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

func (h *odontologoHandler) GetByID(c *gin.Context) {
	id, err := util.GetIdFromParam(c)
	if err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	odontologo, err := h.service.GetByID(id)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if odontologo.Id != id {
		web.NotFound(c)
		return
	}
	web.Success(c, http.StatusOK, odontologo)
}

func (h *odontologoHandler) GetAll(c *gin.Context) {
	dentists, err := h.service.GetAll()
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, dentists)
}
