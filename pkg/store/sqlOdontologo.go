package store

import (
	"database/sql"

	"github.com/rdavid87/sistema-reserva-turnos/internal/domain"
)

type sqlOdontologo struct {
	db *sql.DB
}

func NewSqlOdontologo(db *sql.DB) StoreOdontologo {
	return &sqlOdontologo{
		db: db,
	}
}

func (s *sqlOdontologo) Add(odontologo *domain.Odontologo) (int, error) {
	res, err := s.db.Exec("INSERT INTO odontologos(apellido, nombre, matricula) VALUES (?, ?, ?)", odontologo.Apellido, odontologo.Nombre, odontologo.Matricula)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	odontologo.Id = int(id)

	return odontologo.Id, nil
}

func (s *sqlOdontologo) Update(odontologo *domain.Odontologo) error {
	_, err := s.db.Exec("UPDATE odontologos SET apellido = ?, nombre = ?, matricula = ? WHERE id = ?", odontologo.Apellido, odontologo.Nombre, odontologo.Matricula, odontologo.Id)
	return err
}

func (s *sqlOdontologo) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM odontologos WHERE id = ?", id)
	return err
}

func (s *sqlOdontologo) GetByID(id int) (*domain.Odontologo, error) {
	row := s.db.QueryRow("SELECT id, apellido, nombre, matricula FROM odontologos WHERE id = ?", id)

	var odontologo domain.Odontologo
	err := row.Scan(&odontologo.Id, &odontologo.Apellido, &odontologo.Nombre, &odontologo.Matricula)
	if err != nil {
		if err == sql.ErrNoRows {
			return &odontologo, nil
		}
		return &odontologo, err
	}

	return &odontologo, nil
}

func (s *sqlOdontologo) GetByMatricula(matricula string) (*domain.Odontologo, error) {
	row := s.db.QueryRow("SELECT id, apellido, nombre, matricula FROM odontologos WHERE matricula = ?", matricula)

	var odontologo domain.Odontologo
	err := row.Scan(&odontologo.Id, &odontologo.Apellido, &odontologo.Nombre, &odontologo.Matricula)
	if err != nil {
		if err == sql.ErrNoRows {
			return &odontologo, nil
		}
		return &odontologo, err
	}

	return &odontologo, nil
}

func (s *sqlOdontologo) GetAll() ([]*domain.Odontologo, error) {
	rows, err := s.db.Query("SELECT id, apellido, nombre, matricula FROM odontologos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	odontologos := make([]*domain.Odontologo, 0)
	for rows.Next() {
		var odontologo domain.Odontologo
		err := rows.Scan(&odontologo.Id, &odontologo.Apellido, &odontologo.Nombre, &odontologo.Matricula)
		if err != nil {
			return nil, err
		}
		odontologos = append(odontologos, &odontologo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return odontologos, nil
}
