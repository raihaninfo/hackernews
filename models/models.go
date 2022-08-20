package models

import "github.com/upper/db/v4"

type Model struct {
	User userModel
}

func New(db db.Session) Model {
	return Model{
		User: userModel{db: db},
	}
}
