package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	globalvar "goFileStruc/GlobalVar"
	"goFileStruc/Sales_Analysis/models"
	"goFileStruc/helper"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetUploadFile(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUploadFile(+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")
	if r.Method == http.MethodGet {
		var lResponse models.ResponseStruc
		lResponse.Status = globalvar.Success
		log.Println("Operation Done --------------(+)--")
		FileReader()
		log.Println("Operation Done --------------(-)--")

		lErr := json.NewEncoder(w).Encode(&lResponse)
		if lErr != nil {
			fmt.Fprintf(w, helper.GetErrorString("GUF01", lErr.Error()))
		}

	} else {
		fmt.Fprintf(w, helper.GetErrorString("GUF02", "invalid Method"))
	}

	log.Println("GetUploadFile(-)")

}

func FileReader() {
	log.Println("FileReader(+)")
	var Customer models.CustomerStruc
	var Product models.ProductStruc
	var Order models.OrderStruc

	// Open the CSV file
	file, err := os.Open("csvFile/saleanalytics.csv")
	// file, err := os.Open("csvFile/data.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV: %s", err)
	}

	for _, col := range records[1:] {
		// Read each record and assign to struct
		Order.Order_Id = col[0]
		Product.ProductId = col[1]
		Customer.Customer_Id = col[2]
		Product.ProductName = col[3]
		Product.Category = col[4]
		Order.Region = col[5]
		Order.Date_of_sale = col[6]
		Order.Quantity_sold, _ = strconv.Atoi(col[7])
		Order.Unit_price, _ = strconv.Atoi(col[8])
		Order.Discount, _ = strconv.Atoi(col[9])
		Order.Shipping_cost, _ = strconv.Atoi(col[10])
		Order.Payment_method = col[11]
		Customer.Customer_name = col[12]
		Customer.Customer_email = col[13]
		Customer.Customer_address = col[14]

		Order.Product_Id = col[1]
		Order.Customer_Id = col[2]

		log.Println("Customer: ", Customer)
		log.Println("Product: ", Product)
		log.Println("Order: ", Order)

		// Here you can add code to save the structs to the database
		// db.Create(&Customer)
		// db.Create(&Product)
		// db.Create(&Order)
	}

	log.Println("FileReader(-)")
}
