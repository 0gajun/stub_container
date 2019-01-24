package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	appName := getAppName()
	log.Printf("Running stub container for %s service", appName)

	healthCheckPath := getHealthCheckPath()

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc(healthCheckPath, healthCheckHandler)
	r.HandleFunc("/nettest/{host}/{port:[0-9]+}", netTestHandler)

	addr := getListenAddr()
	if err := http.ListenAndServe(addr, r); err != nil {
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

func getHealthCheckPath() string {
	path, found := os.LookupEnv("HEALTH_CHECK_PATH")
	if !found {
		return "/health_check"
	}
	return path
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	appName := getAppName()
	fmt.Fprintf(w, "stub container for %s service\n", appName)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

func netTestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	target := vars["host"] + ":" + vars["port"]

	timeout := time.Duration(3) * time.Second
	_, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		fmt.Fprintf(w, "Failed to establish tcp connection to %s\n", target)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully established tcp connection to %s\n", target)
}
