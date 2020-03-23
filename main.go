package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
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

func f_addStaticFolders(m martini.ClassicMartini) {
	staticHTML := martini.StaticOptions{Prefix: "html"}
	m.Use(martini.Static("html", staticHTML))

	staticCSS := martini.StaticOptions{Prefix: "css"}
	m.Use(martini.Static("html/css", staticCSS))

	staticIMG := martini.StaticOptions{Prefix: "img"}
	m.Use(martini.Static("html/img", staticIMG))
	staticVUE := martini.StaticOptions{Prefix: "vue"}
	m.Use(martini.Static("html/vue", staticVUE))
	staticJS := martini.StaticOptions{Prefix: "js"}
	m.Use(martini.Static("html/js", staticJS))
	staticFONTS := martini.StaticOptions{Prefix: "fonts"}
	m.Use(martini.Static("html/fonts", staticFONTS))

	staticTPL := martini.StaticOptions{Prefix: "tpl"}
	m.Use(martini.Static("html/tpl", staticTPL))
}

func f_checkErr(err error) {
	if err != nil {
		log.Println("***[ERROR***]", err.Error())
		panic(err.Error())
	}
}

func f_tpl() *template.Template {

	tmpl, err := template.New("").Delims("[[[", "]]]").ParseFiles(
		"html/index.html",
	)

	f_checkErr(err)
	return tmpl
}

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

func getAllData(w http.ResponseWriter) {

	db := dbconnect()
	defer db.Close()
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

func dbconnect() *sql.DB {
	connectonString := username + ":" + password + "@tcp(" + host + ")/" + database + "?parseTime=true"

	db, err := sql.Open("mysql", connectonString)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		log.Println("db connect successfully")
	}

	return db
}

func index(w http.ResponseWriter) {
	f_tpl().ExecuteTemplate(w, "index", nil)
}

func main() {

	m := martini.Classic()
	f_addStaticFolders(*m)

	fmt.Println("run")

	m.Get("/", index)
	m.Get("/data", getAllData)
	m.RunOnAddr(":3000")

}
