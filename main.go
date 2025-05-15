package main

import (
	"goFileStruc/dbconnection"
	"log"
	"os"
	"time"
)

func main() {
	lFile, lErr := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if lErr != nil {
		log.Fatalf("error opening file: %v", lErr)
	}
	defer lFile.Close()
	log.SetOutput(lFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("[ERROR] log file created")

	dbconnection.Dbconnection()

}
