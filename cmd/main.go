package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	serv := server.New(logger)
	err := serv.Serv.ListenAndServe()

	if err != nil {
		logger.Fatal(err)
	}
}
