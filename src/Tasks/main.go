package main

import (
	"log"
	"net/http"
)

func main() {
	PORT := ":8080"
	log.Print("Running on server "+ PORT)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/",CompleteTaskFunc)
	http.HandleFunc("/task/",GetTaskFunc)
	log.Fatal(http.ListenAndServe(PORT,nil))
}

func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(r.URL.Path))
}

func GetTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Path[len("/task/"):]
		_, _ = w.Write([]byte("Get the task "+id))
	}
}