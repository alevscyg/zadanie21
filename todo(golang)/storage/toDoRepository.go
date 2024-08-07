package storage

import (
	"github.com/alevscyg/zadanie23/internal/model"
)

// ToDoRepository ...
type ToDoRepository interface {
	Create(*model.ToDo) (*model.ToDo, error)
	FindById(int) (*model.ToDo, error)
}
