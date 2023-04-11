package paciente

import "github.com/rdavid87/sistema-reserva-turnos/internal/domain"

type Service interface {
	Add(paciente *domain.Paciente) (int, error)
	Update(paciente *domain.Paciente) error
	Delete(id int) error
	GetByID(id int) (*domain.Paciente, error)
	GetAll() ([]*domain.Paciente, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Add(paciente *domain.Paciente) (int, error) {
	return s.repo.Add(paciente)
}

func (s *service) Update(paciente *domain.Paciente) error {
	return s.repo.Update(paciente)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *service) GetByID(id int) (*domain.Paciente, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetAll() ([]*domain.Paciente, error) {
	return s.repo.GetAll()
}
