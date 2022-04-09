package main

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
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
	log.Println("started....")
}
