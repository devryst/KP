package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// employee struct (model)
type Employees struct {
	EmployeeID        string `json:"EmployeeID"`
	LastName          string `json:"LastName"`
	FirstName         string `json:"FirstName"`
	Title             string `json:"Title"`
	TitleOfCourtesy   string `json:"TitleOfCourtesy"`
	BirthDate      	  string `json:"BirthDate"`
	HireDate       	  string `json:"HireDate"`
	Address           string `json:"Address"`
	City 	          string `json:"City"`
	Region     	      string `json:"Region"`
	PostalCode        string `json:"PostalCode"`
	Country     	  string `json:"Country"`
	HomePhone         string `json:"HomePhone"`
	Extension         string `json:"Extension"`
	Photo             string `json:"Photo"`
	Notes       	  string `json:"Notes"`
	ReportsTo         string `json:"ReportsTo"`
	ProvinceName      string `json:"ProvinceName"`
	
}

// get all orders
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Employee []Employees

	sql := `SELECT
				EmployeeID,
				IFNULL(LastName,''),
				IFNULL(FirstName,'') FirstName,
				IFNULL(Title,'') Title,
				IFNULL(TitleOfCourtesy,'') TitleCourtesy,
				IFNULL(BirthDate,'') HireDate,
				IFNULL(HireDate,'') HireDate,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Region,'') Region,
				IFNULL(PostalCode,'') PostalCode,
				IFNULL(Country,'') Country,
				IFNULL(HomePhone,'') HomePhone,
				IFNULL(Extension,'') Extension,
				IFNULL(Photo,'') Photo,
				IFNULL(Notes,'') Notes,
				IFNULL(ReportsTo,'') ReportsTo,
				IFNULL(ProvinceName,'') ProvinceName,
				
			FROM employees`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}
	for result.Next() {

		var employees Employees
		err := result.Scan(&employees.EmployeeID, &employees.LastName, &employees.FirstName, &employees.Title, &employees.TitleOfCourtesy, &employees.BirthDate,
			&employees.HireDate, &employees.Address, &employees.City, &employees.Region, &employees.PostalCode, &employees.Country, &employees.HomePhone,
			&employees.Extension, &employees.Photo, &employees.Notes, &employees.ReportsTo, &employees.ProvinceName)

		if err != nil {
			panic(err.Error())
		}
		employee = append(employee, Employees)
	}
	json.NewEncoder(w).Encode(employee)
}

func createEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		EmployeeID := r.FormValue("EmployeeID")
		LastName := r.FormValue("LastName")
		FirstName := r.FormValue("FirstName")
		Title := r.FormValue("Title")
		TitleOfCourtesy := r.FormValue("TitleOfCourtesy")
		BirthDate := r.FormValue("BirthDate")
		HireDate := r.FormValue("HireDate")
		Address := r.FormValue("Address")
		City := r.FormValue("City")
		Region := r.FormValue("Region")
		PostalCode := r.FormValue("PostalCode")
		Country := r.FormValue("Country")
		HomePhone := r.FormValue("HomePhone")
		Extension := r.FormValue("Extension")
		Photo := r.FormValue("Photo")
		Notes := r.FormValue("Notes")
		ReportsTo := r.FormValue("ReportsTo")
		ProvinceName := r.FormValue("ProvinceName")

		
		stmt, err := db.Prepare("INSERT INTO employee (EmployeeID,LastName,FirstName,Title,TitleOfCourtesy,BirthDate,HireDate,Address,City,Region,PostalCode,Country,HomePhone,Extension,Photo,Notes,ReportsTo,ProvinceName) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

		_, err = stmt.Exec(EmployeeID,LastName,FirstName,Title,TitleOfCourtesy,BirthDate,HireDate,Address,City,Region,PostalCode,Country,HomePhone,Extension,Photo,Notes,ReportsTo,ProvinceName)
		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Created")
		}

	}
}
func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee []employees
	params := mux.Vars(r)

	sql := `SELECT
				EmployeeID,
				IFNULL(LastName,''),
				IFNULL(FirstName,'') FirstName,
				IFNULL(Title,'') Title,
				IFNULL(TitleOfCourtesy,'') TitleCourtesy,
				IFNULL(BirthDate,'') HireDate,
				IFNULL(HireDate,'') HireDate,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Region,'') Region,
				IFNULL(PostalCode,'') PostalCode,
				IFNULL(Country,'') Country,
				IFNULL(HomePhone,'') HomePhone,
				IFNULL(Extension,'') Extension,
				IFNULL(Photo,'') Photo,
				IFNULL(Notes,'') Notes,
				IFNULL(ReportsTo,'') ReportsTo,
				IFNULL(ProvinceName,'') ProvinceName,
				
			FROM employees`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}
	for result.Next() {

		var employees Employees
		err := result.Scan(&employees.EmployeeID, &employees.LastName, &employees.FirstName, &employees.Title, &employees.TitleOfCourtesy, &employees.BirthDate,
			&employees.HireDate, &employees.Address, &employees.City, &employees.Region, &employees.PostalCode, &employees.Country, &employees.HomePhone,
			&employees.Extension, &employees.Photo, &employees.Notes, &employees.ReportsTo, &employees.ProvinceName)

		if err != nil {
			panic(err.Error())
		}
		employee = append(employee, Employees)
	}
	json.NewEncoder(w).Encode(employee)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		EmployeeID := r.FormValue("EmployeeID")
		LastName := r.FormValue("LastName")
		FirstName := r.FormValue("FirstName")
		Title := r.FormValue("Title")
		TitleOfCourtesy := r.FormValue("TitleOfCourtesy")
		BirthDate := r.FormValue("BirthDate")
		HireDate := r.FormValue("HireDate")
		Address := r.FormValue("Address")
		City := r.FormValue("City")
		Region := r.FormValue("Region")
		PostalCode := r.FormValue("PostalCode")
		Country := r.FormValue("Country")
		HomePhone := r.FormValue("HomePhone")
		Extension := r.FormValue("Extension")
		Photo := r.FormValue("Photo")
		Notes := r.FormValue("Notes")
		ReportsTo := r.FormValue("ReportsTo")
		ProvinceName := r.FormValue("ProvinceName")

	stmt, err := db.Prepare("INSERT INTO customers (EmployeeID,LastName,FirstName,Title,TitleOfCourtesy,BirthDate,HireDate,Address,City,Region,PostalCode,Country,HomePhone,Extension,Photo,Notes,ReportsTo,ProvinceName) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var employees employees
	for result.Next() {
		err := result.Scan(&employees.EmployeeID, &employees.LastName, &employees.FirstName, &employees.Title, &employees.TitleOfCourtesy, &employees.BirthDate,
			&employees.HireDate, &employees.Address, &employees.City, &employees.Region, &employees.PostalCode, &employees.Country, &employees.HomePhone,
			&employees.Extension, &employees.Photo, &employees.Notes, &employees.ReportsTo, &employees.ProvinceName)

		if err != nil {
			panic(err.Error())
		}
		employee = append(employee, employees)
	}
	json.NewEncoder(w).Encode(employee)
}
func updateEmployees(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {
		params := mux.Vars(r)
		newLastName := r.FormValue("LastName")
		stmt, err := db.Prepare("UPDATE employee SET LastName = ? WHERE EmployeeID = ?")
		_, err = stmt.Exec(newLastName, params["id"])
		if err != nil {
			fmt.Fprintf(w, " data not found or request error")
		}
		fmt.Fprintf(w, "LastName = %s was deleted", params["id"])
	}
}

func deleteEmployees(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM employee WHERE EmployeeID = ?")

	_, err = stmt.Exec(params["id"])
	if err != nil {
		fmt.Fprintf(w, "delete failed")
	}
	fmt.Fprintf(w, "employee = %s was deleted", params["id"])
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var employee []Employees

	EmployeeID := r.FormValue("EmployeeID")
	LastName := r.FormValue("LastName")

	sql := `SELECT
				EmployeeID,
				IFNULL(LastName,''),
				IFNULL(FirstName,'') FirstName,
				IFNULL(Title,'') Title,
				IFNULL(TitleOfCourtesy,'') TitleCourtesy,
				IFNULL(BirthDate,'') HireDate,
				IFNULL(HireDate,'') HireDate,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Region,'') Region,
				IFNULL(PostalCode,'') PostalCode,
				IFNULL(Country,'') Country,
				IFNULL(HomePhone,'') HomePhone,
				IFNULL(Extension,'') Extension,
				IFNULL(Photo,'') Photo,
				IFNULL(Notes,'') Notes,
				IFNULL(ReportsTo,'') ReportsTo,
				IFNULL(ProvinceName,'') ProvinceName,
	
			FROM employees WHERE EmployeeID = ? AND LastName = ?`

	result, err := db.Query(sql, EmployeeID, LastName)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var employees Employees

	for result.Next() {
		err := result.Scan(&employees.EmployeeID, &employees.LastName, &employees.FirstName, &employees.Title, &employees.TitleOfCourtesy, &employees.BirthDate,
			&employees.HireDate, &employees.Address, &employees.City, &employees.Region, &employees.PostalCode, &employees.Country, &employees.HomePhone,
			&employees.Extension, &employees.Photo, &employees.Notes, &employees.ReportsTo, &employees.ProvinceName)
		if err != nil {
			panic(err.Error())
		}
		employee = append(employee, employees)
	}
	json.NewEncoder(w).Encode(employee)
}

//main function
func main() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_testing")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//init router
	r := mux.NewRouter()

	//route hendles
	r.HandleFunc("/employee", getEmployees).Methods("GET")
	r.HandleFunc("/employee/{id}", getEmployee).Methods("GET")
	r.HandleFunc("/employee", createEmployees).Methods("POST")
	r.HandleFunc("/employee/{id}", updateEmployees).Methods("PUT")
	r.HandleFunc("/employee/{id}", deleteEmployees).Methods("DELETE")

	//new
	r.HandleFunc("/getemployee", getPost).Methods("POST")

	//start server
	log.Fatal(http.ListenAndServe(":8282", r))
}
