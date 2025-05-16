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

	// Open the CSV file
	file, err := os.Open("csvFile/data.csv")
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

	// // Iterate through the records
	// for i, record := range records[1:] {
	// 	// Print header separately
	// 	if i == 0 {
	// 		fmt.Println("Header:", record)
	// 	} else {
	// 		fmt.Printf("Row %d: Name=%s, Age=%s, Location=%s\n", i, record[0], record[1], record[2])
	// 	}
	// }

	for j, value := range records[1:] {
		fmt.Printf("  Column %d: %s\n", j+1, value[0])
	}

	log.Println("FileReader(-)")
}
