package model

type ToDo struct {
	Id          int    `json:"id"`
	Userid      int    `json:"userid"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
