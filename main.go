package main

import "fmt"


const (
	host       = "db.xkc1.ru:27017"
	database   = "crm"
	username   = "crm"
	password   = "Robot611"
	collection = "crm"
	userId     = "e7429bbeab5fa68668e3de2fd2d3484b8e55053ffbd2c1d5bd896ae07067d17e"
)

type toDoList struct {
	Id string
}

func main()  {


	fmt.Println("run")
}



db.createUser( {
	user:"crm",
	pwd: passwordPrompt(),   // Instead of specifying the password in cleartext
	roles:[ "readWrite" ]
} )