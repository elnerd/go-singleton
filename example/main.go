package main

import (
	"github.com/elnerd/go-singleton"
	"github.com/elnerd/go-singleton/example/internal/models"
	"github.com/elnerd/go-singleton/example/internal/server"
	"time"
)

// main initializes the application configuration, stores it globally, starts the server, and waits before exiting.
// the server will retrieve the global configuration and print it on the screen
func main() {
	appConfig := models.AppConfig{
		DBConnStr: "postgres://user:password@localhost:5432/mydb",
		APIKey:    "test-api-key-12345",
	}
	singleton.Store("global-config", &appConfig)

	go server.Start()
	// Sleep for 6 seconds before we quit
	time.Sleep(6000 * time.Millisecond)

}
