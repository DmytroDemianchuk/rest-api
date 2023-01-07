package cmd

import (
	"github.com/dmytrodemianchuk/rest-api/pkg/handler"
	"github.com/dmytrodemianchuk/rest-api/pkg/repository"
	"github.com/dmytrodemianchuk/rest-api/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)x

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
