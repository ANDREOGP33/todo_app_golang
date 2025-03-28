package services

import (
	"todo/internal/application/domain"
	"todo/internal/application/ports"
)

type tarefaService struct {
	tarefaRepo ports.TarefaReposytory
}

func NewTarefaService(tarefaRepo ports.TarefaReposytory) ports.TarefaService {
	return &tarefaService{
		tarefaRepo: tarefaRepo,
	}
}

func (t tarefaService) CreateTarefa(tarefa *domain.Tarefa) error {
	return t.tarefaRepo.Create(tarefa)
}

func (t tarefaService) DeleteTarefa(tarefa *domain.Tarefa) error {
	return t.tarefaRepo.Delete(tarefa)
}

func (t tarefaService) UpdateFeito(tarefa *domain.Tarefa) error {
	return t.tarefaRepo.UpdateFeito(tarefa)
}

func (t tarefaService) ListTarefas() ([]*domain.Tarefa, error) {
	return t.tarefaRepo.List()
}
