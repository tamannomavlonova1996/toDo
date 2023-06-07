package main

import (
	"context"
	"fmt"
	"ginCli"
	"ginCli/handlers"
	"ginCli/repository"
	"ginCli/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s, err := repository.NewStorage()
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}

	newService := service.NewService(s)

	h := handlers.Handler{
		Service: newService,
	}

	srv := new(ginCli.Server)
	go func() {
		if err := srv.Run("8000", h.InitRoutes()); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
			return
		}
	}()

	fmt.Println("Hello, App Started!")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	fmt.Println("Shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Errorf("error occurred on server shutting down: %s", err.Error())
	}
}
