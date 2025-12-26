package main

import (
	"errors"
	"sync"
)

type TodoStore struct {
	mu     sync.RWMutex
	Todos  []Todo
	nextID int
}

func NewTodoStore() *TodoStore {
	return &TodoStore{
		Todos:  []Todo{},
		nextID: 1,
	}
}

func (s *TodoStore) DeleteTodoByID(ID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, todo := range s.Todos {
		if todo.ID == ID {
			s.Todos = append(s.Todos[:i], s.Todos[i+1:]...)
			return nil
		}
	}

	return errors.New("ID not found for deletion")
}

func (s *TodoStore) UpdateTodoByID(id int, changes map[string]interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var target *Todo

	for i := range s.Todos {
		if s.Todos[i].ID == id {
			target = &s.Todos[i]
			break
		}
	}

	if target == nil {
		return errors.New("Invalid ID")
	}

	for k, v := range changes {
		switch k {
		case "Completed":
			completed, ok := v.(bool)
			if !ok {
				return errors.New("Completed must be true or false")
			}
			target.Completed = completed

		case "Description":
			description, ok := v.(string)
			if !ok {
				return errors.New("Invalid description")
			}
			target.Description = description
		}
	}

	return nil
}

func (s *TodoStore) FindTodoByID(id int) (*Todo, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for i := range s.Todos {
		if s.Todos[i].ID == id {
			return &s.Todos[i], true
		}
	}

	return nil, false
}

func (s *TodoStore) GetTodos() []Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]Todo, len(s.Todos))
	copy(result, s.Todos)
	return result
}

func (s *TodoStore) AddTodo(description string) Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	newTodo := Todo{
		ID:          s.nextID,
		Description: description,
		Completed:   false,
	}

	s.Todos = append(s.Todos, newTodo)
	s.nextID++
	return newTodo
}
