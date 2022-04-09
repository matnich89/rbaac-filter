package main

import (
	"github.com/gorilla/mux"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
)

func main() {
	log.Println("starting.....")
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalln("This application has to be ran instead a k8s cluster...")
	}
	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("issue getting new config %s", err.Error())
	}
	http.ListenAndServe(":8080", mux.NewRouter())
}
