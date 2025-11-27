package entities

import (
	"prauth/services"

	"github.com/gorilla/sessions"
)

type AppCtx struct{
	CookieStore sessions.CookieStore
}

type User struct {
	ID    int `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	PwdHash string `json:"pwdhash"`
	Dbs *services.DataBaseService
}

func (u *User) Create() error {
	err := u.Dbs.Conn.QueryRow(u.Dbs.Ctx,
		"INSERT INTO users (name, email, pwdhash) VALUES ($1,$2,$3) RETURNING id",
		u.Name, u.Email, u.PwdHash).Scan(&u.ID)
	return err
}

func (u *User) GetByEmail() error {
	err := u.Dbs.Conn.QueryRow(u.Dbs.Ctx, "SELECT id, name, email FROM users WHERE email=$1", u.Email).
	Scan(&u.ID, &u.Name, &u.Email)
	return err
}

func (u *User) Update() error{ 
	_, err := u.Dbs.Conn.Exec(
		u.Dbs.Ctx,
		"UPDATE users SET name=$1, email=$2 WHERE id=$3",
		u.Name, u.Email, u.ID,
	)
	return err
}

func (u *User) Delete() error {
	_, err := u.Dbs.Conn.Exec(u.Dbs.Ctx, "DELETE FROM users WHERE id=$1", u.ID)
	return err
}
