package model

import "time"

type User struct {
	UserId    int       `json:"userId" db:"user_id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	CreatedOn time.Time `json:"createdOn" db:"created_on"`
	Lastlogin time.Time `json:"lastlogin" db:"last_login"`
}
