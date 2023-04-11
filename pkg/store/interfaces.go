package store

import "github.com/rdavid87/sistema-reserva-turnos/internal/domain"

type StoreOdontologo interface {
	Add(odontologo *domain.Odontologo) (int, error)
	Update(odontologo *domain.Odontologo) error
	Delete(id int) error
	GetByID(id int) (*domain.Odontologo, error)
	GetAll() ([]*domain.Odontologo, error)
}

type StorePaciente interface {
	Add(paciente *domain.Paciente) (int, error)
	Update(paciente *domain.Paciente) error
	Delete(id int) error
	GetByID(id int) (*domain.Paciente, error)
	GetAll() ([]*domain.Paciente, error)
}
