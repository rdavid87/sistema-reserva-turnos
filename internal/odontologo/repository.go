package odontologo

import (
	"github.com/rdavid87/sistema-reserva-turnos/internal/domain"
	"github.com/rdavid87/sistema-reserva-turnos/pkg/store"
)

type Repository interface {
	Add(odontologo *domain.Odontologo) (int, error)
	Update(odontologo *domain.Odontologo) error
	Delete(id int) error
	GetByID(id int) (*domain.Odontologo, error)
	GetByMatricula(matricula string) (*domain.Odontologo, error)
	GetAll() ([]*domain.Odontologo, error)
}

type repository struct {
	storage store.StoreOdontologo
}

func NewRepository(storage store.StoreOdontologo) Repository {
	return &repository{storage}
}

func (r *repository) Add(odontologo *domain.Odontologo) (int, error) {
	return r.storage.Add(odontologo)
}

func (r *repository) Update(odontologo *domain.Odontologo) error {
	return r.storage.Update(odontologo)
}

func (r *repository) Delete(id int) error {
	return r.storage.Delete(id)
}

func (r *repository) GetByID(id int) (*domain.Odontologo, error) {
	return r.storage.GetByID(id)
}

func (r *repository) GetByMatricula(matricula string) (*domain.Odontologo, error) {
	return r.storage.GetByMatricula(matricula)
}

func (r *repository) GetAll() ([]*domain.Odontologo, error) {
	return r.storage.GetAll()
}
