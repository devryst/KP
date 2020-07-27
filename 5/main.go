package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type Root struct {
	XMLName   xml.Name `xml:"Root"`
	Text      string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	Customers struct {
		Text     string `xml:",chardata"`
		Customer []struct {
			Text         string `xml:",chardata"`
			CustomerID   string `xml:"CustomerID,attr"`
			CompanyName  string `xml:"CompanyName"`
			ContactName  string `xml:"ContactName"`
			ContactTitle string `xml:"ContactTitle"`
			Phone        string `xml:"Phone"`
			FullAddress  struct {
				Text       string `xml:",chardata"`
				Address    string `xml:"Address"`
				City       string `xml:"City"`
				Region     string `xml:"Region"`
				PostalCode string `xml:"PostalCode"`
				Country    string `xml:"Country"`
			} `xml:"FullAddress"`
			Fax string `xml:"Fax"`
		} `xml:"Customer"`
	} `xml:"Customers"`
	Orders struct {
		Text  string `xml:",chardata"`
		Order []struct {
			Text         string `xml:",chardata"`
			CustomerID   string `xml:"CustomerID"`
			EmployeeID   string `xml:"EmployeeID"`
			OrderDate    string `xml:"OrderDate"`
			RequiredDate string `xml:"RequiredDate"`
			ShipInfo     struct {
				Text           string `xml:",chardata"`
				ShippedDate    string `xml:"ShippedDate,attr"`
				ShipVia        string `xml:"ShipVia"`
				Freight        string `xml:"Freight"`
				ShipName       string `xml:"ShipName"`
				ShipAddress    string `xml:"ShipAddress"`
				ShipCity       string `xml:"ShipCity"`
				ShipRegion     string `xml:"ShipRegion"`
				ShipPostalCode string `xml:"ShipPostalCode"`
				ShipCountry    string `xml:"ShipCountry"`
			} `xml:"ShipInfo"`
		} `xml:"Order"`
	} `xml:"Orders"`
}

func getCustomers(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Root

	if err = xml.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	//Tugas insert kan ke table Customer
	for i := 0; i < len(request.Customers.Customer); i++ {

		CustomerID := request.Customers.Customer[i].CustomerID
		CompanyName := request.Customers.Customer[i].CompanyName
		ContactName := request.Customers.Customer[i].ContactName
		ContactTitle := request.Customers.Customer[i].ContactTitle
		Phone := request.Customers.Customer[i].Phone
		Address := request.Customers.Customer[i].FullAddress.Address
		City := request.Customers.Customer[i].FullAddress.City
		Region := request.Customers.Customer[i].FullAddress.Region
		PostalCode := request.Customers.Customer[i].FullAddress.PostalCode
		Country := request.Customers.Customer[i].FullAddress.Country
		Fax := request.Customers.Customer[i].Fax

		stmt, err := db.Prepare("INSERT INTO Customers (CustomerID,CompanyName,ContactName,ContactTitle,Phone,Address,City,Region,PostalCode,Country,Fax) Values (?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		_, err = stmt.Exec(CustomerID, CompanyName, ContactName, ContactTitle, Phone, Address, City, Region, PostalCode, Country, Fax)
	}

	//Tugas insert kan ke table Order
	for i := 0; i < len(request.Customers.Customer); i++ {

		CustomerID := request.Orders.Order[i].CustomerID
		EmployeeID := request.Orders.Order[i].EmployeeID
		OrderDate := request.Orders.Order[i].OrderDate
		RequiredDate := request.Orders.Order[i].RequiredDate
		ShippedDate := request.Orders.Order[i].ShipInfo.ShippedDate
		ShipVia := request.Orders.Order[i].ShipInfo.ShipVia
		Freight := request.Orders.Order[i].ShipInfo.Freight
		ShipName := request.Orders.Order[i].ShipInfo.ShipName
		ShipAddress := request.Orders.Order[i].ShipInfo.ShipAddress
		ShipCity := request.Orders.Order[i].ShipInfo.ShipCity
		ShipRegion := request.Orders.Order[i].ShipInfo.ShipRegion
		ShipPostalCode := request.Orders.Order[i].ShipInfo.ShipPostalCode
		ShipCountry := request.Orders.Order[i].ShipInfo.ShipCountry

		stmt, err := db.Prepare("INSERT INTO Orders (CustomerID,EmployeeID,OrderDate,RequiredDate,ShippedDate,ShipVia,Freight,ShipName,ShipAddress,ShipCity,ShipRegion,ShipPostalCode,ShipCountry) Values (?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		_, err = stmt.Exec(CustomerID, EmployeeID, OrderDate, RequiredDate, ShippedDate, ShipVia, Freight, ShipName, ShipAddress, ShipCity, ShipRegion, ShipPostalCode, ShipCountry)
	}

}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/northwind")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	fmt.Println("Server on :8181")

	// Route handles & endpoints
	r.HandleFunc("/customers", getCustomers).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8181", r))

}
