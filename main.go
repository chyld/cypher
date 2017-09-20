package main

import (
	"database/sql"
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

var db *sql.DB

func main() {
	server()
}

func server() {
	e := echo.New()
	e.Debug = true
	e.GET("/", home)
	e.GET("/logins", index)
	e.POST("/logins", create)
	e.Logger.Fatal(e.Start(":3333"))
}

func connect() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:pass1234@localhost/temp?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func query(sql string) *sql.Rows {
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	return rows
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func index(c echo.Context) error {
	connect()
	rows := query("SELECT * FROM logins")
	defer rows.Close()
	defer db.Close()

	ls := make([]*Login, 0)
	for rows.Next() {
		l := new(Login)
		rows.Scan(&l.ID, &l.Email, &l.Username, &l.Password, &l.Pin, &l.Site, &l.Meta, &l.CreatedAt)
		ls = append(ls, l)
	}
	return c.JSON(http.StatusOK, ls)
}

func create(c echo.Context) error {
	connect()
	defer db.Close()

	l := new(Login)
	if err := c.Bind(l); err != nil {
		return c.String(http.StatusBadRequest, "Could not bind JSON to struct")
	}

	_, err := db.Exec("INSERT INTO logins (email, username, password, pin, site, meta, created_at) VALUES($1, $2, $3, $4, $5, $6, $7)", l.Email, l.Username, l.Password, l.Pin, l.Site, l.Meta, time.Now())
	if err != nil {
		return c.String(http.StatusBadRequest, "Could not insert record into database")
	}

	return c.JSON(http.StatusOK, "Insert successful")
}
