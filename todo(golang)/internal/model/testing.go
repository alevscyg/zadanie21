package model

import "testing"

// TestToDo ...
func TestToDo(t *testing.T) *ToDo {
	t.Helper()

	return &ToDo{
		Userid:      1,
		Title:       "Заварить чай",
		Description: "...",
	}
}
