package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

var (
	dhcpdConfigFile string
	dhcpdLeaseFile  string
	port            string
)

func GetStateEndpoint(w http.ResponseWriter, req *http.Request) {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "dhcpd-pools"
	cmdArgs := []string{"--config",
		dhcpdConfigFile,
		"--leases",
		dhcpdLeaseFile,
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
	dhcpdConfigFile = os.Getenv("DHCPD_CONF_FILE")
	dhcpdLeaseFile = os.Getenv("DHCPD_LEASE_FILE")
	port = os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/state", GetStateEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
