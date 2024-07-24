package model_test

import (
	"testing"

	"github.com/alevscyg/zadanie23/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestToDo(t)
	assert.NotEmpty(t, u)
}
