package store

import (
	"database/sql"

	"github.com/rdavid87/sistema-reserva-turnos/internal/domain"
)

type sqlPaciente struct {
	db *sql.DB
}

func NewSqlPaciente(db *sql.DB) StorePaciente {
	return &sqlPaciente{
		db: db,
	}
}

func (s *sqlPaciente) Add(paciente *domain.Paciente) (int, error) {
	res, err := s.db.Exec("INSERT INTO pacientes(apellido, nombre, dni, domicilio, fecha_alta) VALUES (?, ?, ?, ?, ?)", paciente.Apellido, paciente.Nombre, paciente.DNI, paciente.Domicilio, paciente.FechaAlta)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	paciente.Id = int(id)

	return paciente.Id, nil
}

func (s *sqlPaciente) Update(paciente *domain.Paciente) error {
	_, err := s.db.Exec("UPDATE pacientes SET apellido = ?, nombre = ?, dni = ?, domicilio = ?, fecha_alta = ? WHERE id = ?", paciente.Apellido, paciente.Nombre, paciente.DNI, paciente.Domicilio, paciente.FechaAlta, paciente.Id)
	return err
}

func (s *sqlPaciente) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM pacientes WHERE id = ?", id)
	return err
}

func (s *sqlPaciente) GetByID(id int) (*domain.Paciente, error) {
	row := s.db.QueryRow("SELECT id, apellido, nombre, dni, domicilio, fecha_alta FROM pacientes WHERE id = ?", id)

	var paciente domain.Paciente
	err := row.Scan(&paciente.Id, &paciente.Apellido, &paciente.Nombre, &paciente.DNI, &paciente.Domicilio, &paciente.FechaAlta)
	if err != nil {
		if err == sql.ErrNoRows {
			return &paciente, nil
		}
		return &paciente, err
	}

	return &paciente, nil
}

func (s *sqlPaciente) GetAll() ([]*domain.Paciente, error) {
	rows, err := s.db.Query("SELECT id, apellido, nombre, dni, domicilio, fecha_alta FROM pacientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pacientes := make([]*domain.Paciente, 0)
	for rows.Next() {
		var paciente domain.Paciente
		err := rows.Scan(&paciente.Id, &paciente.Apellido, &paciente.Nombre, &paciente.DNI, &paciente.Domicilio, &paciente.FechaAlta)
		if err != nil {
			return nil, err
		}
		pacientes = append(pacientes, &paciente)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pacientes, nil
}
