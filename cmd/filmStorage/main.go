package main

import (
	"film_storage/pkg/configs"
	"film_storage/pkg/logging"
	"fmt"
	"net/http"

	"log/slog"
	"os"
)

func main() {

	os.Setenv("CONFIG_PATH", "/Users/mirustal/Documents/project/go/film_storage/config.yml")
	os.Setenv("SECRET_KEY", "Medods_Task1")

	cfg := configs.GetConfig()
	fmt.Println("hello this ", cfg)
	log := logging.SetupLogger(cfg.ModeLog)
	log.Info("Starting service", slog.String("env", cfg.ModeLog))

	s := &http.Server{
		Addr: cfg.Http.Port,
		// Handler:        myHandler,
	}
}
