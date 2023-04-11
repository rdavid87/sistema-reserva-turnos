package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rdavid87/sistema-reserva-turnos/cmd/server/handler"
	"github.com/rdavid87/sistema-reserva-turnos/internal/odontologo"
	"github.com/rdavid87/sistema-reserva-turnos/internal/paciente"
	"github.com/rdavid87/sistema-reserva-turnos/pkg/store"
)

func main() {
	// Conexión a la base de datos
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sistema_turnos")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	storageOdontologo := store.NewSqlOdontologo(db)
	odontologoRepo := odontologo.NewRepository(storageOdontologo)
	odontologoService := odontologo.NewService(odontologoRepo)
	odontologoHandler := handler.NewOdontologoHandler(odontologoService)

	storagePaciente := store.NewSqlPaciente(db)
	pacienteRepo := paciente.NewRepository(storagePaciente)
	pacienteService := paciente.NewService(pacienteRepo)
	pacienteHandler := handler.NewPacienteHandler(pacienteService)

	// Creación del enrutador HTTP
	r := gin.Default()

	// Rutas para el CRUD de dentistas
	odologologoRoutes := r.Group("/odontologo")
	{
		odologologoRoutes.POST("/", odontologoHandler.Add)
		odologologoRoutes.GET("/:id", odontologoHandler.GetByID)
		odologologoRoutes.GET("/", odontologoHandler.GetAll)
		odologologoRoutes.PUT("/:id", odontologoHandler.Update)
		odologologoRoutes.PATCH("/:id", odontologoHandler.Patch)
		odologologoRoutes.DELETE("/:id", odontologoHandler.Delete)
	}

	// Rutas para el CRUD de pacientes
	pacienteRoutes := r.Group("/paciente")
	{
		pacienteRoutes.POST("/", pacienteHandler.Add)
		pacienteRoutes.GET("/:id", pacienteHandler.GetByID)
		pacienteRoutes.GET("/", pacienteHandler.GetAll)
		pacienteRoutes.PUT("/:id", pacienteHandler.Update)
		pacienteRoutes.PATCH("/:id", pacienteHandler.Patch)
		pacienteRoutes.DELETE("/:id", pacienteHandler.Delete)
	}

	// Inicio del servidor HTTP
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
