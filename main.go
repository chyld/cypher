package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

type login struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Pin       string    `json:"pin"`
	Site      string    `json:"site"`
	Meta      string    `json:"meta"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	e := echo.New()
	e.GET("/logins", func(c echo.Context) error {
		return c.String(http.StatusOK, "logins lolz 4")
	})
	e.Logger.Fatal(e.Start(":3333"))

	db, err := sql.Open("postgres", "postgres://postgres:pass1234@localhost/temp?sslmode=disable")
	fmt.Printf("db: %#v, err: %#v", db, err)
	rows, err := db.Query("SELECT * FROM logins")
	fmt.Printf("rows: %#v, err: %#v", rows, err)
	for rows.Next() {
		l := login{}
		err := rows.Scan(&l.ID, &l.Email, &l.Username, &l.Password, &l.Pin, &l.Site, &l.Meta, &l.CreatedAt)
		fmt.Printf("\n\n\nerr: %#v, login: %#v", err, l)
		y, m, d := l.CreatedAt.Date()
		fmt.Printf("\n\n\ny: %#v, m: %#v, d: %#v", y, m, d)
	}
}
