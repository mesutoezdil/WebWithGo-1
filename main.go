package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", Homepage)
	http.ListenAndServe(":8080", nil)
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("index.html")
	data := make(map[string]interface{})
	data["Numbers"] = []int{1, 2, 3, 4, 5}
	data["is admin"] = false
	data["numbers"] = 10
	view.Execute(w, data)
}
