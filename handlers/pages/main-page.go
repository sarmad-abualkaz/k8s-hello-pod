package pages

import (
	"net/http"
	"html/template"
	"k8s-hello-pod/config"
)

// Return the main page from html template:
func MainPage(w http.ResponseWriter, r *http.Request) {
	
	pod := config.ConfigEnv.Pod

	t := template.Must(template.ParseFiles("templates/main.html"))
	
	t.Execute(w, pod)
}