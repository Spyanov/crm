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
	DealTitle   string    `json:"dealTitle"`
	DealDesc    string    `json:"dealDesc"`
	Price       int       `json:"price"`
	StartPeriod time.Time `json:"startPeriod"`
	EndPeriod   time.Time `json:"endPeriod"`
	Status      string    `json:"status"`
	Resul       string    `json:"resul"`
}

type page struct {
	TitlePage string     `json:"titlePage"`
	StatusArr []string   `json:"statusArr"`
	Status1   []toDoList `json:"status_1"`
	Status2   []toDoList `json:"status_2"`
	Status3   []toDoList `json:"status_3"`
	Status4   []toDoList `json:"status_4"`
	Status5   []toDoList `json:"status_5"`
	Status6   []toDoList `json:"status_6"`
	Status7   []toDoList `json:"status_7"`
}

func getAllData(w http.ResponseWriter) {

	db := dbconnect()
	defer db.Close()

	var cell toDoList
	var status1, status2, status3, status4, status5, status6 []toDoList

	// status 1
	rows, err := db.Query("SELECT * FROM todolist WHERE status='Потенциальная идея'")
	if err != nil {
		fmt.Println("[dbconnect] Ошибка запроса в БД", err)
	}

	for rows.Next() {
		err = rows.Scan(&cell.Id, &cell.Visible, &cell.Client, &cell.DealTitle, &cell.DealDesc, &cell.Price, &cell.StartPeriod, &cell.EndPeriod, &cell.Status, &cell.Resul)
		if err != nil {
			fmt.Println("Ошибка сканирования строки результата", err)
		}
		status1 = append(status1, cell)
	}
	// status 1

	// status 2
	rows, err = db.Query("SELECT * FROM todolist WHERE status='Предложение'")
	if err != nil {
		fmt.Println("[dbconnect] Ошибка запроса в БД", err)
	}

	for rows.Next() {
		err = rows.Scan(&cell.Id, &cell.Visible, &cell.Client, &cell.DealTitle, &cell.DealDesc, &cell.Price, &cell.StartPeriod, &cell.EndPeriod, &cell.Status, &cell.Resul)
		if err != nil {
			fmt.Println("Ошибка сканирования строки результата", err)
		}
		status2 = append(status2, cell)
	}
	// status 2

	// status 3
	rows, err = db.Query("SELECT * FROM todolist WHERE status='Согласование'")
	if err != nil {
		fmt.Println("[dbconnect] Ошибка запроса в БД", err)
	}

	for rows.Next() {
		err = rows.Scan(&cell.Id, &cell.Visible, &cell.Client, &cell.DealTitle, &cell.DealDesc, &cell.Price, &cell.StartPeriod, &cell.EndPeriod, &cell.Status, &cell.Resul)
		if err != nil {
			fmt.Println("Ошибка сканирования строки результата", err)
		}
		status3 = append(status3, cell)
	}
	// status 3

	// status 3
	rows, err = db.Query("SELECT * FROM todolist WHERE status='Согласование, подготовка'")
	if err != nil {
		fmt.Println("[dbconnect] Ошибка запроса в БД", err)
	}

	for rows.Next() {
		err = rows.Scan(&cell.Id, &cell.Visible, &cell.Client, &cell.DealTitle, &cell.DealDesc, &cell.Price, &cell.StartPeriod, &cell.EndPeriod, &cell.Status, &cell.Resul)
		if err != nil {
			fmt.Println("Ошибка сканирования строки результата", err)
		}
		status3 = append(status3, cell)
	}
	// status 3

	// status 4
	rows, err = db.Query("SELECT * FROM todolist WHERE status='В работе'")
	if err != nil {
		fmt.Println("[dbconnect] Ошибка запроса в БД", err)
	}

	for rows.Next() {
		err = rows.Scan(&cell.Id, &cell.Visible, &cell.Client, &cell.DealTitle, &cell.DealDesc, &cell.Price, &cell.StartPeriod, &cell.EndPeriod, &cell.Status, &cell.Resul)
		if err != nil {
			fmt.Println("Ошибка сканирования строки результата", err)
		}
		status4 = append(status4, cell)
	}
	// status 4

	// status 5
	rows, err = db.Query("SELECT * FROM todolist WHERE status='Закрытие'")
	if err != nil {
		fmt.Println("[dbconnect] Ошибка запроса в БД", err)
	}

	for rows.Next() {
		err = rows.Scan(&cell.Id, &cell.Visible, &cell.Client, &cell.DealTitle, &cell.DealDesc, &cell.Price, &cell.StartPeriod, &cell.EndPeriod, &cell.Status, &cell.Resul)
		if err != nil {
			fmt.Println("Ошибка сканирования строки результата", err)
		}
		status5 = append(status5, cell)
	}
	// status 5

	// status 6
	rows, err = db.Query("SELECT * FROM todolist WHERE status='Ждем оплаты'")
	if err != nil {
		fmt.Println("[dbconnect] Ошибка запроса в БД", err)
	}

	for rows.Next() {
		err = rows.Scan(&cell.Id, &cell.Visible, &cell.Client, &cell.DealTitle, &cell.DealDesc, &cell.Price, &cell.StartPeriod, &cell.EndPeriod, &cell.Status, &cell.Resul)
		if err != nil {
			fmt.Println("Ошибка сканирования строки результата", err)
		}
		status6 = append(status6, cell)
	}
	// status 6

	var globalPage page

	globalPage.Status1 = status1
	globalPage.Status2 = status2
	globalPage.Status3 = status3
	globalPage.Status4 = status4
	globalPage.Status5 = status5
	globalPage.Status6 = status6
	//globalPage.Status7 = status7

	toJson, err := json.Marshal(globalPage)
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
