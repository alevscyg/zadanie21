package sqlstoragetest_test

import (
	"testing"

	"github.com/alevscyg/zadanie23/internal/model"
	sqlstoragetest "github.com/alevscyg/zadanie23/storage/sqlStorage_test"
	"github.com/stretchr/testify/assert"
)

func TestToDoRepository_Create(t *testing.T) {
	s := sqlstoragetest.New()
	u := model.TestToDo(t)
	result, err := s.ToDo().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, result.Id)
}

func TestToDoRepository_FindById(t *testing.T) {
	s := sqlstoragetest.New()
	u1 := model.TestToDo(t)
	s.ToDo().Create(u1)
	u2, err := s.ToDo().FindById(u1.Userid)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
