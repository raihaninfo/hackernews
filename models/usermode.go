package models

import "github.com/upper/db/v4"

type userModel struct {
	db db.Session
}

func (m userModel) get() {
	m.db.Collection("user").Find()
}
