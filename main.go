package main

import (
	"gofr.dev/cmd/gofr/migration"
	dbmigration "gofr.dev/cmd/gofr/migration/dbMigration"
	"gofr.dev/pkg/gofr"

	"sample/handler"
	"sample/migrations"
	"sample/store"
)

func main() {
	// Creating GoFr app
	app := gofr.New()

	// Running migrations - UP
	if err := migration.Migrate("remote-config-data", dbmigration.NewGorm(app.GORM()),
		migrations.All(), dbmigration.UP, app.Logger); err != nil {
		app.Logger.Fatalf("Error in running migrations: %v", err)
	}

	carStore := store.New()
	carHandler := handler.New(carStore)

	// Creating routes
	app.POST("/car", carHandler.Create)
	app.GET("/cars", carHandler.GetAll)
	app.GET("/car/{id}", carHandler.GetByID)
	app.PUT("/car/{id}", carHandler.Update)
	app.DELETE("/car/{id}", carHandler.Delete)

	// Starting server
	app.Start()
}
