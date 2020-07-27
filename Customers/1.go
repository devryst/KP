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

// Customer struct (Model) ...
type Customer struct {
	CustomerID   string `json:"CustomerID"`
	CompanyName  string `json:"CompanyName"`
	ContactName  string `json:"ContactName"`
	ContactTitle string `json:"ContactTitle"`
	Address      string `json:"Address"`
	City         string `json:"City"`
	Country      string `json:"Country"`
	Phone        string `json:"Phone"`
	PostalCode   string `json:"PostalCode"`
}

// Get all orders

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var customers []Customer

	sql := `SELECT
				CustomerID,
				IFNULL(CompanyName,''),
				IFNULL(ContactName,'') ContactName,
				IFNULL(ContactTitle,'') ContactTitle,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Country,'') Country,
				IFNULL(Phone,'') Phone ,
				IFNULL(PostalCode,'') PostalCode
			FROM customers`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var customer Customer
		err := result.Scan(&customer.CustomerID, &customer.CompanyName, &customer.ContactName,
			&customer.ContactTitle, &customer.Address, &customer.City, &customer.Country,
			&customer.Phone, &customer.PostalCode)

		if err != nil {
			panic(err.Error())
		}
		customers = append(customers, customer)
	}

	json.NewEncoder(w).Encode(customers)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		CustomerID := r.FormValue("CustomerID")
		CompanyName := r.FormValue("CompanyName")
		ContactName := r.FormValue("ContactName")
		ContactTitle := r.FormValue("ContactTitle")
		Address := r.FormValue("Address")
		City := r.FormValue("City")
		Region := r.FormValue("Region")
		PostalCode := r.FormValue("PostalCode")
		Country := r.FormValue("Country")
		Phone := r.FormValue("Phone")
		Fax := r.FormValue("Fax")

		stmt, err := db.Prepare("INSERT INTO customers (CustomerID,CompanyName,ContactName,ContactTitle,Address,City,Region,PostalCode,Country,Phone,Fax) VALUES (?,?,?,?,?,?,?,?,?,?,?)")

		_, err = stmt.Exec(CustomerID, CompanyName, ContactName, ContactTitle, Address, City, Region, PostalCode, Country, Phone, Fax)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Created")
		}

	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []Customer
	params := mux.Vars(r)

	sql := `SELECT
				CustomerID,
				IFNULL(CompanyName,''),
				IFNULL(ContactName,'') ContactName,
				IFNULL(ContactTitle,'') ContactTitle,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Country,'') Country,
				IFNULL(Phone,'') Phone ,
				IFNULL(PostalCode,'') PostalCode
			FROM customers WHERE CustomerID = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var customer Customer

	for result.Next() {

		err := result.Scan(&customer.CustomerID, &customer.CompanyName, &customer.ContactName,
			&customer.ContactTitle, &customer.Address, &customer.City, &customer.Country,
			&customer.Phone, &customer.PostalCode)

		if err != nil {
			panic(err.Error())
		}

		customers = append(customers, customer)
	}

	json.NewEncoder(w).Encode(customers)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newCompanyName := r.FormValue("CompanyName")

		stmt, err := db.Prepare("UPDATE customers SET CompanyName = ? WHERE CustomerID = ?")

		_, err = stmt.Exec(newCompanyName, params["id"])

		if err != nil {
			fmt.Fprintf(w, "Data not found or Request error")
		}

		fmt.Fprintf(w, "Customer with CustomerID = %s was updated", params["id"])
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM customers WHERE CustomerID = ?")

	_, err = stmt.Exec(params["id"])

	if err != nil {
		fmt.Fprintf(w, "delete failed")
	}

	fmt.Fprintf(w, "Customer with ID = %s was deleted", params["id"])
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var customers []Customer

	CustomerID := r.FormValue("CustomerID")
	CompanyName := r.FormValue("CompanyName")

	sql := `SELECT
				CustomerID,
				IFNULL(CompanyName,''),
				IFNULL(ContactName,'') ContactName,
				IFNULL(ContactTitle,'') ContactTitle,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Country,'') Country,
				IFNULL(Phone,'') Phone ,
				IFNULL(PostalCode,'') PostalCode
			FROM customers WHERE CustomerID = ? AND CompanyName = ?`

	result, err := db.Query(sql, CustomerID, CompanyName)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var customer Customer

	for result.Next() {

		err := result.Scan(&customer.CustomerID, &customer.CompanyName, &customer.ContactName,
			&customer.ContactTitle, &customer.Address, &customer.City, &customer.Country,
			&customer.Phone, &customer.PostalCode)

		if err != nil {
			panic(err.Error())
		}

		customers = append(customers, customer)
	}

	json.NewEncoder(w).Encode(customers)

}

// Main function
func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_testing")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/customers", getCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	r.HandleFunc("/customers", createCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	//New
	r.HandleFunc("/getcustomer", getPost).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
