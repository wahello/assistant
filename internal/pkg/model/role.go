package model

import "time"

type Role struct {
	ID         int    `db:"id"`
	Profession string `db:"profession"`

	Exp   int `db:"exp"`
	Level int `db:"level"`

	// attr
	Strength    int `db:"strength"`
	Culture     int `db:"culture"`
	Environment int `db:"environment"`
	Charisma    int `db:"charisma"`
	Talent      int `db:"talent"`
	Intellect   int `db:"intellect"`

	// -> Equipments
	// -> Quests

	Time time.Time `db:"time"`
}

type Equipment struct {
	ID       int       `db:"id"`
	Name     string    `db:"name"`
	Quality  string    `db:"quality"`
	Level    int       `db:"level"`
	Category string    `db:"category"`
	Time     time.Time `db:"time"`
}

type Quest struct {
	ID            int       `db:"id"`
	Title         string    `db:"title"`
	Exp           int       `db:"exp"`
	AttrPoints    string    `db:"attr_points"`
	Preconditions string    `db:"preconditions"`
	Time          time.Time `db:"time"`
}
