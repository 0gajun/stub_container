package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	appName := getAppName()
	log.Printf("Running stub container for %s service", appName)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "stub container for %s service", appName)
	})

	addr := getListenAddr()
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("unexpected error : %v", err)
	}
}
func getListenAddr() string {
	port, found := os.LookupEnv("LISTEN_PORT")
	if !found {
		return "0.0.0.0:3000"
	}

	return "0.0.0.0:" + port
}
func getAppName() string {
	name, found := os.LookupEnv("APP_NAME")
	if !found {
		return "undefined"
	}
	return name
}
