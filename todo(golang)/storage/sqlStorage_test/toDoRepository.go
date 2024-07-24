package sqlstoragetest

import (
	"github.com/alevscyg/zadanie23/internal/model"
	"github.com/alevscyg/zadanie23/storage"
)

// ToDoRepository ...
type ToDoRepository struct {
	store *Storage
	toDo  map[int]*model.ToDo
}

// Create ...
func (r *ToDoRepository) Create(u *model.ToDo) (*model.ToDo, error) {
	u.Id = len(r.toDo) + 1
	r.toDo[u.Id] = u
	if u.Userid == 0 || len(u.Title) < 1 {
		return nil, storage.ErrRecordNotFound
	}
	return u, nil
}

// FindById ...
func (r *ToDoRepository) FindById(id int) (*model.ToDo, error) {
	u, ok := r.toDo[id]
	if !ok {
		return nil, storage.ErrRecordNotFound
	}

	return u, nil
}
