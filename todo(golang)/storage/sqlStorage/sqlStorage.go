package sqlstorage

import (
	"database/sql"

	"github.com/alevscyg/zadanie23/storage"
	_ "github.com/lib/pq" // ...
)

// Storage ...
type Storage struct {
	db             *sql.DB
	ToDoRepository *ToDoRepository
}

// New ...
func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

// ToDo ...
func (s *Storage) ToDo() storage.ToDoRepository {
	if s.ToDoRepository != nil {
		return s.ToDoRepository
	}

	return &ToDoRepository{
		storage: s,
	}
}
