package main

import (
	"github.com/elnerd/go-singleton"
	"github.com/elnerd/go-singleton/example/internal/models"
	"github.com/elnerd/go-singleton/example/internal/server"
)

func main() {
	appConfig := models.AppConfig{
		DBConnStr: "postgres://user:password@localhost:5432/mydb",
		APIKey:    "test-api-key-12345",
	}
	singleton.Store("global-config", &appConfig)

	go server.Start()
}
