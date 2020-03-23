package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

const (
	host     = "db.xkc1.ru:3306"
	database = "crm"
	username = "admin"
	password = "NUjFcwmP!666999"
)

type toDoList struct {
	Id          int       `json:"id"`
	Visible     bool      `json:"visible"`
	Client      string    `json:"client"`
	DealTitle   string    `json:"deal_title"`
	DealDesc    string    `json:"deal_desc"`
	Price       int       `json:"price"`
	StartPeriod time.Time `json:"start_period"`
	EndPeriod   time.Time `json:"end_period"`
	Status      string    `json:"status"`
	Resul       string    `json:"resul"`
}

func dbconnect(w http.ResponseWriter) {
	connectonString := username + ":" + password + "@tcp(" + host + ")/" + database + "?parseTime=true"

	db, err := sql.Open("mysql", connectonString)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		log.Println("db connect successfully")
	}

	rows, err := db.Query("SELECT * FROM todolist")
	if err != nil {
		fmt.Println("[dbconnect] Ошибка запроса в БД", err)
	}

	var cell toDoList
	var arr []toDoList

	for rows.Next() {
		err = rows.Scan(&cell.Id, &cell.Visible, &cell.Client, &cell.DealTitle, &cell.DealDesc, &cell.Price, &cell.StartPeriod, &cell.EndPeriod, &cell.Status, &cell.Resul)
		if err != nil {
			fmt.Println("Ошибка сканирования строки результата", err)
		}
		arr = append(arr, cell)
	}

	toJson, err := json.Marshal(arr)
	if err != nil {
		fmt.Println("Ошибка конвертации в Json")
	}

	w.Write(toJson)
}

func main() {

	m := martini.Classic()

	fmt.Println("run")

	m.Get("/", dbconnect)
	m.RunOnAddr(":8000")

}
