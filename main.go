package main

import (
	"database/sql"
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
	catalogRepo := repo.ConfigureMySQLServiceRepo(client)
	serviceversionRepo := repo.ConfigureMySQLServiceVersionRepo(client)
	catalogSvc := service.ConfigureServiceImpl(catalogRepo, serviceversionRepo)
	v := validator.New()
	routes.InitServiceHandler(e, catalogSvc, v)
	e.Start(":1323")
}

func ConnectDB() (*ent.Client, error) {
	dsn := "root:cashfree.123@tcp(localhost:3306)/test?parseTime=true"
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
