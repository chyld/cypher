package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type dog struct {
	Name string `json:"alias"`
	Age  int64  `json:"old"`
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:pass1234@localhost/temp")
	fmt.Printf("db: %#v, err: %#v", db, err)

	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

// HelloServer is awesome
func HelloServer(w http.ResponseWriter, req *http.Request) {
	go hs(w, req)
}

func hs(w http.ResponseWriter, req *http.Request) {
	d := dog{"fido", 3}
	fmt.Printf("d is : %#v", d)
	// a := []string{"apple", "peach", "pear"}
	b, _ := json.Marshal(d)
	fmt.Printf("b is : %#v", b)
	time.Sleep(5000 * time.Millisecond)
	io.WriteString(w, string(b))
}
