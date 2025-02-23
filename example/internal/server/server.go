package server

import (
	"fmt"
	"github.com/elnerd/go-singleton/example/internal/models"
	"github.com/elnerd/go-singleton/pkg/singleton"
	"time"
)

var appConfig *models.AppConfig

func Start() {
	fmt.Println("Starting server")
	fmt.Println("Retrieving appConfig from singleton...")
	if err := singleton.GetInto("global-config", &appConfig); err != nil {
		panic(err)
	}
	fmt.Println("Retrieved appConfig from singleton. APIKey is successfully retrieved.")

	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Stopping server")
}
