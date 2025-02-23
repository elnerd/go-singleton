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
	if err := singleton.GetInto("global-config", &appConfig); err != nil {
		panic(err)
	}

	for {
		fmt.Println("Server is running")
		fmt.Println(appConfig)
		time.Sleep(1000)
	}
}
