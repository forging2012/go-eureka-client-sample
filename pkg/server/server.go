package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/carllhw/go-eureka-client-sample/pkg/eureka"
	"github.com/carllhw/go-eureka-client-sample/pkg/service"
)

func Start() {
	handleSigterm() // Graceful shutdown on Ctrl+C or kill

	go startWebServer() // Starts HTTP service  (async)

	eureka.Register() // Performs Eureka registration

	go eureka.StartHeartbeat() // Performs Eureka heartbeating (async)

	// Block...
	wg := sync.WaitGroup{} // Use a WaitGroup to block main() exit
	wg.Add(1)
	wg.Wait()
}

func handleSigterm() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		eureka.Deregister()
		os.Exit(1)
	}()
}

func startWebServer() {
	router := service.NewRouter()
	log.Println("Starting HTTP service at 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port 8080")
		log.Println("Error: " + err.Error())
	}
}
