package ports

import (
	"todo/internal/application/domain"
)

type TarefaService interface {
	CreateTarefa(tarefa *domain.Tarefa) error
	DeleteTarefa(tarefa *domain.Tarefa) error
	UpdateFeito(tarefa *domain.Tarefa) error
	ListTarefas() ([]*domain.Tarefa, error)
}

type TarefaReposytory interface {
	Create(tarefa *domain.Tarefa) error
	Delete(tarefa *domain.Tarefa) error
	UpdateFeito(tarefa *domain.Tarefa) error
	List() ([]*domain.Tarefa, error)
}
