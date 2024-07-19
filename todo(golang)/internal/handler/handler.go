package handler

import (
	"github.com/alevscyg/zadanie23/internal/model"
	"github.com/alevscyg/zadanie23/storage"
)

func CreateToDo(userid int, title string, description string, storage storage.Storage) (*model.ToDo, error) {
	u := &model.ToDo{
		Userid:      userid,
		Title:       title,
		Description: description,
	}
	todo, err := storage.ToDo().Create(u)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func FindToDoById(userId int, storage storage.Storage) (*model.ToDo, error) {
	data, err := storage.ToDo().FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}
