package repositories

import (
	"errors"
	"log"
	"sync"
	"todo/internal/application/domain"
)

type tarefaRepository struct {
	tarefa map[string]*domain.Tarefa
	mu     sync.RWMutex
}

func NewTarefaRepository() *tarefaRepository {
	return &tarefaRepository{
		tarefa: make(map[string]*domain.Tarefa),
	}
}

func (r *tarefaRepository) Create(tarefa *domain.Tarefa) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	tarefa.Concluida = false
	r.tarefa[tarefa.Nome] = tarefa
	return nil
}

func (r *tarefaRepository) UpdateFeito(tarefa *domain.Tarefa) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	tarefa, exists := r.tarefa[tarefa.Nome]
	if !exists {
		return errors.New("tarefa não encontrada")
	}
	tarefa.Concluida = !tarefa.Concluida
	r.tarefa[tarefa.Nome] = tarefa
	log.Printf("Tarefa: %v atualizada\n", tarefa.Nome)
	return nil
}

func (r *tarefaRepository) Delete(tarefa *domain.Tarefa) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.tarefa[tarefa.Nome]
	if !exists {
		return errors.New("tarefa não encontrada")
	}
	delete(r.tarefa, tarefa.Nome)
	log.Printf("Tarefa: %v deletada\n", tarefa.Nome)
	return nil
}

func (r *tarefaRepository) List() ([]*domain.Tarefa, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	tarefas := make([]*domain.Tarefa, 0)
	for _, tarefa := range r.tarefa {
		tarefas = append(tarefas, tarefa)
	}
	return tarefas, nil
}
