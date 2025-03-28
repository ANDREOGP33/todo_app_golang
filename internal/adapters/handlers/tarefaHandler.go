package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"todo/internal/application/domain"
	"todo/internal/application/ports"
	"todo/internal/util"
)

type TarefaHandler struct {
	tarefaService ports.TarefaService
}

func NewTarefaHandler(tarefaService ports.TarefaService) *TarefaHandler {
	return &TarefaHandler{
		tarefaService: tarefaService,
	}
}

func (h *TarefaHandler) Init(c *gin.Context) {
	tarefas, _ := h.tarefaService.ListTarefas()
	util.ExecutarTemplate(c, "index.html", gin.H{
		"Tarefas": tarefas,
	})
}

func (h *TarefaHandler) CreateTarefa(c *gin.Context) {
	var tarefaRequest struct {
		Nome      string `form:"nome" json:"nome"`
		Descricao string `form:"descricao" json:"descricao"`
	}

	if err := c.ShouldBind(&tarefaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tarefa := &domain.Tarefa{
		Nome:        tarefaRequest.Nome,
		Descricao:   tarefaRequest.Descricao,
		Concluida:   false,
		DataCriacao: time.Now().Format("01/02/2006 15:04:05"),
	}

	if err := h.tarefaService.CreateTarefa(tarefa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tarefa criada com sucesso"})
}

func (h *TarefaHandler) GetAllTarefas(c *gin.Context) {
	tarefas, _ := h.tarefaService.ListTarefas()
	c.JSON(http.StatusOK, gin.H{
		"tarefas": tarefas,
	})
}

func (h *TarefaHandler) DeleteTarefa(c *gin.Context) {
	var tarefaRequest struct {
		Nome string `form:"nome" json:"nome"`
	}

	if err := c.ShouldBind(&tarefaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tarefa := &domain.Tarefa{
		Nome: tarefaRequest.Nome,
	}

	if err := h.tarefaService.DeleteTarefa(tarefa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarefa deletada com sucesso"})
}

func (h *TarefaHandler) UpdateTarefa(c *gin.Context) {
	var tarefaRequest struct {
		Nome string `form:"nome" json:"nome"`
	}

	if err := c.ShouldBind(&tarefaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tarefa := &domain.Tarefa{
		Nome: tarefaRequest.Nome,
	}

	if err := h.tarefaService.UpdateFeito(tarefa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarefa atualizada com sucesso"})
}
