package domain

import db "lla/db/sqlc"

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (user *User) FromEntity(dbUser *db.User) {
	user.ID = dbUser.ID
	user.Email = dbUser.Email
	user.FirstName = dbUser.FirstName
	user.LastName = dbUser.LastName
}
