package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

// Login struct
type Login struct {
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
	server()
}

func server() {
	e := echo.New()
	e.GET("/", home)
	e.GET("/logins", logins)
	e.Logger.Fatal(e.Start(":3333"))
}

func connect() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:pass1234@localhost/temp?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func query(db *sql.DB, sql string) *sql.Rows {
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	return rows
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func logins(c echo.Context) error {
	db := connect()
	rows := query(db, "SELECT * FROM logins")
	defer rows.Close()
	defer db.Close()

	ls := make([]*Login, 0)
	for rows.Next() {
		l := new(Login)
		rows.Scan(&l.ID, &l.Email, &l.Username, &l.Password, &l.Pin, &l.Site, &l.Meta, &l.CreatedAt)
		ls = append(ls, l)
	}
	fmt.Printf("ls: %#v", ls)

	return c.JSON(http.StatusOK, ls)
}
