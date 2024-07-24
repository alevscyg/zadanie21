package sqlstorage

import (
	"database/sql"

	"github.com/alevscyg/zadanie23/internal/model"
	store "github.com/alevscyg/zadanie23/storage"
)

// ToDoRepository ...
type ToDoRepository struct {
	storage *Storage
}

// Create ...
func (r *ToDoRepository) Create(u *model.ToDo) (*model.ToDo, error) {
	if err := r.storage.db.QueryRow(
		"INSERT INTO todo (userid, title, description) VALUES ($1, $2, $3) RETURNING id",
		u.Userid,
		u.Title,
		u.Description,
	).Scan(&u.Id); err != nil {
		return nil, err
	}
	return u, nil
}

// FindById ...
func (r *ToDoRepository) FindById(id int) (*model.ToDo, error) {
	u := &model.ToDo{}
	if err := r.storage.db.QueryRow(
		"SELECT id, userid, title, description FROM todo WHERE id = $1",
		id,
	).Scan(
		&u.Id,
		&u.Userid,
		&u.Title,
		&u.Description,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// FindAll ...
func (r *ToDoRepository) FindAll(id int) *sql.Row {
	return r.storage.db.QueryRow(
		"SELECT id, userid, title, description FROM todo")
}
