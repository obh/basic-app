package main

import (
	"database/sql"
	"os"
	"services/ent"
	"services/repo"
	"services/routes"
	"services/service"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	client, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	servicesRepo := repo.ConfigureMySQLServicesRepo(client)
	serviceVersionsRepo := repo.ConfigureMySQLServiceVersionsRepo(client)
	servicesController := service.ConfigureServiceController(servicesRepo, serviceVersionsRepo)
	v := validator.New()
	routes.InitServiceHandler(e, servicesController, v)
	e.Start(":1323")
}

func ConnectDB() (*ent.Client, error) {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}
