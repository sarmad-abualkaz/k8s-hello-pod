package main

import (
	"net/http"
	"k8s-hello-pod/config"
	"k8s-hello-pod/handlers/pages"
	"k8s-hello-pod/handlers/status"
)

func main() {

	port := config.ConfigEnv.Port
	
	http.HandleFunc("/", pages.MainPage)
	
	http.HandleFunc("/status", status.GiveHealth)

	http.ListenAndServe(":"+port, nil)
}
