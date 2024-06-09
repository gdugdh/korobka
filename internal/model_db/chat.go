package model_db

import (
	"time"
)

type Message struct {
	Id       int64     `db:"id"`
	IdAuthor int64     `db:"id_author"`
	IdChat   int64     `db:"id_chat"`
	Content  string    `db:"content"`
	Datetime time.Time `db:"datetime"`
}
