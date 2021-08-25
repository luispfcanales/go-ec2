package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "luis"
	dbname   = "cloud"
)

type Response struct {
	Status      int         `json:"status"`
	TypeMessage string      `json:"type_message"`
	Data        interface{} `json:"data"`
}
type Person struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
type ListPerson []Person

func main() {
	http.HandleFunc("/", Home)
	fmt.Println("server run port:80")
	http.ListenAndServe(":80", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	list := ListPerson{}
	res := &Response{
		Status:      http.StatusOK,
		TypeMessage: "Success",
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	rows, err := db.Query("SELECT username,email FROM cloud")
	CheckError(err)
	defer rows.Close()

	for rows.Next() {
		p := Person{}
		rows.Scan(&p.Username, &p.Email)
		list = append(list, p)
	}

	res.Data = list

	json.NewEncoder(w).Encode(res)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
