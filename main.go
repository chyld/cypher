package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
	var x int
	db, err := sql.Open("postgres", "postgres://postgres:pass1234@localhost/temp")
	fmt.Printf("db: %#v, err: %#v", db, err)
}
