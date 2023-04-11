package paciente

import (
	"github.com/rdavid87/sistema-reserva-turnos/internal/domain"
	"github.com/rdavid87/sistema-reserva-turnos/pkg/store"
)

type Repository interface {
	Add(paciente *domain.Paciente) (int, error)
	Update(paciente *domain.Paciente) error
	Delete(id int) error
	GetByID(id int) (*domain.Paciente, error)
	GetAll() ([]*domain.Paciente, error)
}

type repository struct {
	storage store.StorePaciente
}

func NewRepository(storage store.StorePaciente) Repository {
	return &repository{storage}
}

func (r *repository) Add(paciente *domain.Paciente) (int, error) {
	return r.storage.Add(paciente)
}

func (r *repository) Update(paciente *domain.Paciente) error {
	return r.storage.Update(paciente)
}

func (r *repository) Delete(id int) error {
	return r.storage.Delete(id)
}

func (r *repository) GetByID(id int) (*domain.Paciente, error) {
	return r.storage.GetByID(id)
}

func (r *repository) GetAll() ([]*domain.Paciente, error) {
	return r.storage.GetAll()
}
