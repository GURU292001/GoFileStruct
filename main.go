package main

import (
	"fmt"
	"goFileStruc/Sales_Analysis/handlers"
	"goFileStruc/dbconnection"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	lFile, lErr := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if lErr != nil {
		log.Fatalf("error opening file: %v", lErr)
	}
	defer lFile.Close()
	log.SetOutput(lFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Database connection process
	lErr = dbconnection.BuildConnection()
	if lErr != nil {
		log.Fatal(lErr)
	}
	defer dbconnection.G_Db_instance.Mysql_sqldb.Close()

	// http.HandleFunc("/uploadFile", apis.GetUploadFile)
	fmt.Println("Server started---(+)")
	// http.ListenAndServe(":29069", nil)

	lRouter := mux.NewRouter()
	lRouter.HandleFunc("/uploadFile", handlers.GetUploadFile).Methods(http.MethodGet)

	lSrv := &http.Server{
		Handler: lRouter,
		Addr:    ":29069",
	}

	log.Fatal(lSrv.ListenAndServe())

}
