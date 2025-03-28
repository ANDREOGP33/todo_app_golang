package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
	"todo/internal/adapters/handlers"
	"todo/internal/adapters/repositories"
	"todo/internal/application/services"
	"todo/internal/util"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/static", filepath.Join(".", "web", "static"))

	util.CarregarTemplates(r)

	tarefaRepo := repositories.NewTarefaRepository()
	tarefaService := services.NewTarefaService(tarefaRepo)
	tarefaHandler := handlers.NewTarefaHandler(tarefaService)

	r.GET("/", tarefaHandler.Init)
	r.POST("/criar", tarefaHandler.CreateTarefa)
	r.GET("/tarefas", tarefaHandler.GetAllTarefas)
	r.POST("/delete", tarefaHandler.DeleteTarefa)
	r.POST("/update", tarefaHandler.UpdateTarefa)

	port := ":8080"
	fmt.Println("Listening on port", port)
	log.Fatal(r.Run(port))
}
