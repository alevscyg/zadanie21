package sqlstorage_test

import (
	"testing"

	"github.com/alevscyg/zadanie23/internal/model"
	store "github.com/alevscyg/zadanie23/storage"
	sqlstorage "github.com/alevscyg/zadanie23/storage/sqlStorage"
	"github.com/stretchr/testify/assert"
)

func TestToDoRepository_Create(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("todo")

	s := sqlstorage.New(db)
	u := model.TestToDo(t)

	result, err := s.ToDo().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestToDoRepository_FindById(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("todo")

	s := sqlstorage.New(db)
	todo := model.TestToDo(t)

	_, err := s.ToDo().FindById(todo.Id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	result, err := s.ToDo().Create(todo)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	u2, err := s.ToDo().FindById(result.Id)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
