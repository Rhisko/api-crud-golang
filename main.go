package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"http"
	"log"
	"mux"
)

var (
	db  *gorm.DB
	err error
)

/**Type Data Struct ===========

 */
type Product struct {
	ID    int     `json:"id" `
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Result Array Of Product
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {

	db, err = gorm.Open("mysql", "indosoft:KiT4B15A2018@/go_api_crud?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Println("Connection Failed", err)
	} else {
		fmt.Println("Succes")
	}

	db.AutoMigrate(&Product{})

	handleRequests()

}

func handleRequests() {
	log.Println("Start Server at http://localhost:8080")
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome TO Rest APi")
}
