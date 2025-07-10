// Package internal provides a factory for initializing and managing application components.
package internal

import (
	"github.com/uchupx/worker-api/internal/api/handlers"
	"github.com/uchupx/worker-api/internal/config"
	"github.com/uchupx/worker-api/internal/database"
)

type Factory struct {
	userHandler *handlers.UserHandler

	dbConn *database.MongoDB
}

func (f *Factory) GetUserHandler() *handlers.UserHandler {
	if f.userHandler != nil {
		return f.userHandler
	}

	f.userHandler = &handlers.UserHandler{}

	return f.userHandler
}

func (f *Factory) GetDBConnection() *database.MongoDB {
	if f.dbConn != nil {
		return f.dbConn
	}

	conf := config.GetConfig()

	var err error

	f.dbConn, err = database.NewMongoDB(conf.Database.URL)
	if err != nil {
		panic(err)
	}

	return f.dbConn
}
