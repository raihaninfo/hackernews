package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/upper/db/v4"
)

var (
	ErrNoMoreRows     = errors.New("no record found")
	ErrDuplicateEmail = errors.New("email already in our database")
	ErrUserNotActive = errors.New("your account in inactive")
	ErrInvalidLogin = errors.New("Invalid Login")
)

type Model struct {
	User UserModel
}

func New(db db.Session) Model {
	return Model{
		User: UserModel{db: db},
	}
}

func errHasDuplicate(err error, key string) bool {
	str := fmt.Sprintf(`ERROR: duplicate key value violates unique constraint "%s"`, key)
	return strings.Contains(err.Error(), str)
}

func convertUpperToInt(id db.ID) int {
	idType := fmt.Sprintf("%T", id)
	if idType == "int64" {
		return int(id.(int64))
	}
	return id.(int)
}
