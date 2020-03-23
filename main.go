package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	host       = "db.xkc1.ru:3306"
	database   = "crm"
	username   = "admin"
	password   = "NUjFcwmP!666999"
	collection = "crm"
	userId     = "e7429bbeab5fa68668e3de2fd2d3484b8e55053ffbd2c1d5bd896ae07067d17e"
)

type toDoList struct {
	Id          string `json:"id"`
	Visible     string `json:"visible"`
	Client      string `json:"client"`
	DealTitle   string `json:"deal_title"`
	DealDesc    string `json:"deal_desc"`
	Price       string `json:"price"`
	StartPeriod string `json:"start_period"`
	EndPeriod   string `json:"end_period"`
	Status      string `json:"status"`
	Resul       string `json:"resul"`
}

func dbconnect() {
	connectonString := username + ":" + password + "@tcp(" + host + ")/" + database

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

	for _, v := range arr {
		fmt.Println("arr = ", v)
	}

}

func main() {

	dbconnect()

	fmt.Println("run")
}
