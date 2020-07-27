package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type Foto struct {
	Colors struct {
		Category string `json:"category"`
		Code     struct {
			Hex  string `json:"hex"`
			Rgba string `json:"rgba"`
		} `json:"code"`
		Color string `json:"color"`
		Type  string `json:"type"`
	} `json:"colors"`
	Thumbnail struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"thumbnail"`
}
type Human struct {
	Age    int64 `json:"age"`
	Errors struct {
		Detail string `json:"detail"`
		Source struct {
			Pointer string `json:"pointer"`
		} `json:"source"`
		Status string `json:"status"`
		Title  string `json:"title"`
	} `json:"errors"`
	Name   string `json:"name"`
	Powers struct {
		Task1 string `json:"Task1"`
		Task2 string `json:"Task2"`
		Task3 string `json:"Task3"`
	} `json:"powers"`
	SecretIdentity string `json:"secretIdentity"`
}

func getFoto(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var req Foto

	if err = json.Unmarshal(body, &req); err != nil {
		fmt.Println("Failed decoding json message")
	}

	Category := req.Colors.Category
	Hex := req.Colors.Code.Hex
	Rgba := req.Colors.Code.Rgba
	Color := req.Colors.Color
	Type := req.Colors.Type
	Height := req.Thumbnail.Height
	URL := req.Thumbnail.URL
	Width := req.Thumbnail.Width

	stmt, err := db.Prepare("INSERT INTO foto (Category,Hex,Rgba,Color,Type,Height,URL,Width) Values (?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(Category, Hex, Rgba, Color, Type, Height, URL, Width)
}

func gethuman(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var req Human

	if err = json.Unmarshal(body, &req); err != nil {
		fmt.Println("Failed decoding json message")
	}

	Age := req.Age
	Detail := req.Errors.Detail
	Pointer := req.Errors.Source.Pointer
	Status := req.Errors.Status
	Title := req.Errors.Title
	Name := req.Name
	Task1 := req.Powers.Task1
	Task2 := req.Powers.Task2
	Task3 := req.Powers.Task3
	SecretIdentity := req.SecretIdentity

	stmt, err := db.Prepare("INSERT INTO human (Age,Detail,Pointer,Status,Title,Name,Task1,Task2,Task3,SecretIdentity) Values (?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(Age, Detail, Pointer, Status, Title, Name, Task1, Task2, Task3, SecretIdentity)
}
func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Kp")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	fmt.Println("Server on :8585")

	// Route handles & endpoints
	r.HandleFunc("/foto", getFoto).Methods("POST")
	r.HandleFunc("/human", gethuman).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8585", r))

}
