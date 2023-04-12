package turno

import "github.com/rdavid87/sistema-reserva-turnos/internal/domain"

type Service interface {
	Add(turno *domain.TurnoAbstract) (int, error)
	Update(turno *domain.Turno) error
	Delete(id int) error
	GetByID(id int) (*domain.Turno, error)
	GetAll() ([]*domain.Turno, error)
	GetByDNI(dni string) (*domain.TurnoDetalle, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Add(turno *domain.TurnoAbstract) (int, error) {
	return s.repo.Add(turno)
}

func (s *service) Update(turno *domain.Turno) error {
	return s.repo.Update(turno)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *service) GetByID(id int) (*domain.Turno, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetAll() ([]*domain.Turno, error) {
	return s.repo.GetAll()
}

func (s *service) GetByDNI(dni string) (*domain.TurnoDetalle, error) {
	return s.repo.GetByDNI(dni)
}
