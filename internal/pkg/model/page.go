package model

import "time"

type Page struct {
	ID      int       `db:"id"`
	UUID    string    `db:"uuid"`
	Title   string    `db:"title"`
	Content string    `db:"content"`
	Time    time.Time `db:"time"`
}
