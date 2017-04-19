package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	dhcpdConfigFile = flag.String("dhcpd-config-file", "/etc/dhcpd.conf", "Path to dhcpd configuration file")
	dhcpdLeaseFile  = flag.String("dhcpd-lease-file", "/etc/dhcpd.leases", "Path to dhcpd leases file")
	port            = flag.String("port", "80", "Port to bind the service")
)

func GetStateEndpoint(w http.ResponseWriter, req *http.Request) {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "dhcpd-pools"
	cmdArgs := []string{"--config",
		*dhcpdConfigFile,
		"--leases",
		*dhcpdLeaseFile,
		"--format=j"}
	w.Header().Set("Content-Type", "application/json")
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput(); err != nil {
		w.WriteHeader(500)
		output := fmt.Sprintf("%s, %s", err.Error(), cmdOut)
		log.Println(output)
		json.NewEncoder(w).Encode(map[string]string{"error": output})

	} else {
		out := string(cmdOut)
		w.Write([]byte(out))
	}

}

func main() {
	flag.Parse()
	if _, err := os.Stat(*dhcpdConfigFile); os.IsNotExist(err) {
		log.Fatalf("DHCPD config not found at: %s", *dhcpdConfigFile)
		os.Exit(1)
	}
	if _, err := os.Stat(*dhcpdLeaseFile); os.IsNotExist(err) {
		log.Fatalf("DHCPD leases db not found at: %s", *dhcpdLeaseFile)
		os.Exit(1)
	}
	router := mux.NewRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	router.HandleFunc("/v1/api/state", GetStateEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), loggedRouter))
}
