package odontologo

import "github.com/rdavid87/sistema-reserva-turnos/internal/domain"

type Service interface {
	Add(odontologo *domain.Odontologo) (int, error)
	Update(odontologo *domain.Odontologo) error
	Delete(id int) error
	GetByID(id int) (*domain.Odontologo, error)
	GetAll() ([]*domain.Odontologo, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Add(odontologo *domain.Odontologo) (int, error) {
	return s.repo.Add(odontologo)
}

func (s *service) Update(odontologo *domain.Odontologo) error {
	return s.repo.Update(odontologo)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *service) GetByID(id int) (*domain.Odontologo, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetAll() ([]*domain.Odontologo, error) {
	return s.repo.GetAll()
}
