package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	io.WriteString(w, currentTime.String())
	//log.Println("hahahaahha")
	file, _ := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	log.SetOutput(file)
	log.SetPrefix("[Error]")
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	log.Println("Hi file 这是一个错误")
}

func main() {
	http.HandleFunc("/health", HealthCheck)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
