package main

import (
	"context"
	"fmt"
	"github.com/bagahulho/wbL2/develop/task_11"
	"github.com/bagahulho/wbL2/develop/task_11/pkg/handler"
	"github.com/bagahulho/wbL2/develop/task_11/pkg/repository"
	"github.com/bagahulho/wbL2/develop/task_11/pkg/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(calendar_api.Server)

	go func() {
		if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	fmt.Println("server started on Port 8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("APP shutting down")
	if err := srv.ShutDown(context.Background()); err != nil {
		fmt.Printf("error occurred on server shutting down: %s", err)
	}
}
