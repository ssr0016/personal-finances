package model

import "time"

type User struct {
	Id        string     `db:"id"`
	CreateAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
}
