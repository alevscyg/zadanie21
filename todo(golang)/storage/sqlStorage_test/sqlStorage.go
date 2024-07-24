package sqlstoragetest

import (
	"github.com/alevscyg/zadanie23/internal/model"
	"github.com/alevscyg/zadanie23/storage"
)

// Storage ...
type Storage struct {
	toDoRepository *ToDoRepository
}

// New ...
func New() *Storage {
	return &Storage{}
}

// User ...
func (s *Storage) ToDo() storage.ToDoRepository {
	if s.toDoRepository != nil {
		return s.toDoRepository
	}

	s.toDoRepository = &ToDoRepository{
		store: s,
		toDo:  make(map[int]*model.ToDo),
	}

	return s.toDoRepository
}
