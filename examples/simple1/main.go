package main


import "database/sql"

import (
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)



var db *sql.DB
var gdb *gorm.DB


//get the gorm database
func Database() *gorm.DB {
	gdb, err := gorm.Open("mysql", "root:BENM@tcp(192.168.99.100:30627)/ben")

	if err != nil {
		log.Fatal(err)
	}

	return gdb
}



func testConnect() {
	var (
		id int
		amount float32
		err error
	)



	rows, err := db.Query("select * from ben.transactions")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &amount)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(id, amount)
	}
}


type Users struct {
	Id int
	Firstname string
	Lastname string
}

func GetUsers(c *gin.Context) {
	var users = []Users{
		{Id: 1, Firstname: "Oliver", Lastname: "Queen"},
		{Id: 2, Firstname: "Malcom", Lastname: "Merlyn"},
	}

	c.JSON(200, users)


}

//testing GORM....seems like you need to let GORM create the database

type Transaction struct {
	gorm.Model
	id int `json:"id"`
	amount float32 `json:"amount"`
}

func GetTransactions(c *gin.Context) {
	var transactions []Transaction
	// SELECT * FROM users
	gdb.Find(&transactions)

	// Display JSON result
	c.JSON(200, transactions)

}





type Transaction2 struct {
	Id int
	Amount float32
}


func GetTransactions2(c *gin.Context) {

	var t []Transaction2
	var id int
	var  amount float32

	rows, err := db.Query("select * from ben.transactions")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &amount)
		if err != nil {
			log.Fatal(err)
		}

		t = append(t, Transaction2{id, amount})
	}

	c.JSON(200, t)


}





func main() {
	var err error

	db, err = sql.Open("mysql", "root:BENM@tcp(192.168.99.100:30627)/ben")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	testConnect()
	gdb = Database()

	r := gin.Default()
	r.GET("/USERS", GetUsers)
	r.GET("/TRANS", GetTransactions)
	r.GET("/TRANS2", GetTransactions2)

	r.Run(":8081")


}


