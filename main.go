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

type columnArray struct {
	ColumnTitle   string     `json:"columnTitle"`
	ArrayInColumn []toDoList `json:"arrayInColumn"`
}

type page struct {
	TitlePage      string        `json:"titlePage"`
	DataArr        []columnArray `json:"dataArr"`
	StatusRegistry []string      `json:"statusRegistry"`
	ResultRegistry []string      `json:"resultRegistry"`
}

func getAllData(w http.ResponseWriter) {

	db := dbconnect()
	defer db.Close()

	var allDataPage page

	var trash int
	var currentStatus string
	var StatusRegistryArray []string
	var ResultRegistryArray []string

	// status arr
	rowsList, err := db.Query("SELECT * FROM status ORDER BY id")
	if err != nil {
		fmt.Println("[dbconnect] Ошибка запроса в БД", err)
	}

	for rowsList.Next() {
		var currentColumnArray columnArray

		err = rowsList.Scan(&trash, &currentStatus)
		if err != nil {
			fmt.Println("Ошибка сканирования строки результата", err)
		}

		// запрос на выборку всех данных по текущему статусу
		rowsArrayInColumn, err := db.Query("SELECT * FROM todolist WHERE status ='" + currentStatus + "'")
		if err != nil {
			log.Println("ошибка выборки списка задач из БД по критерию = ", currentStatus, ", ошибка: ", err)
		}
		var cell toDoList
		var arrToDoListInCurrentCell []toDoList

		for rowsArrayInColumn.Next() {
			err = rowsArrayInColumn.Scan(&cell.Id, &cell.Visible, &cell.Client, &cell.DealTitle, &cell.DealDesc, &cell.Price, &cell.StartPeriod, &cell.EndPeriod, &cell.Status, &cell.Resul)
			arrToDoListInCurrentCell = append(arrToDoListInCurrentCell, cell)
		}
		currentColumnArray.ColumnTitle = currentStatus
		currentColumnArray.ArrayInColumn = arrToDoListInCurrentCell
		// .запрос на выборку всех данных по текущему статусу

		StatusRegistryArray = append(StatusRegistryArray, currentStatus)
		allDataPage.DataArr = append(allDataPage.DataArr, currentColumnArray)
	}

	rowsResultRegistry, err := db.Query("SELECT * FROM result ORDER BY id")
	if err != nil {
		fmt.Println("Ошибка запроса статусов акрытия из БД", err)
	}
	var ResultRegistry string
	for rowsResultRegistry.Next() {
		err = rowsResultRegistry.Scan(&trash, &ResultRegistry)
		if err != nil {
			fmt.Println("Ошибка сканирования строки статусов закрытия", err)
		}
		ResultRegistryArray = append(ResultRegistryArray, ResultRegistry)
	}

	allDataPage.StatusRegistry = StatusRegistryArray
	allDataPage.ResultRegistry = ResultRegistryArray

	toJson, err := json.Marshal(allDataPage)
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
