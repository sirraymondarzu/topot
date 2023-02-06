package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there! We are tapot."))
}

func displayTime(w http.ResponseWriter, r *http.Request) {
	localtime := time.Now().Format("3:04:05 PM")
	//read in the template file
	ts, _ := template.ParseFiles("./ui/html/display.time.tmpl")
	ts.Execute(w, localtime)

}
func getValues(w http.ResponseWriter, r *http.Request) {
	

}
func calculateArea(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	mux := http.NewServeMux()

	//register routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/time", displayTime)
	mux.HandleFunc("/area-calculator", getValues)
	mux.HandleFunc("/area-calculator-2", calculateArea)



	//create a file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//create url mapping to the static directory

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	log.Println("Starting on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
