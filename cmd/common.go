package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// LoE : log with error code 1 and print if err is notnull
func LoE(msg string, err error) {
	if err != nil {
		log.Printf("\n❌  %s\n   %v\n", msg, err)
	}
}

// EoE : exit with error code 1 and print, if err is not nil
func EoE(msg string, err error) {
	if err != nil {
		fmt.Printf("\n❌  %s\n   %v\n", msg, err)
		os.Exit(1)
		panic(err)
	}
}

// GetIP : get local ip address
func getPubIP() string {
	// we are using a pulib IP API, we're using ipify here, below are some others
	// https://www.ipify.org
	// http://myexternalip.com
	// http://api.ident.me
	// http://whatismyipaddress.com/api
	// https://ifconfig.co
	// https://ifconfig.me
	url := "https://api.ipify.org?format=text"
	resp, err := http.Get(url)
	EoE("Error Getting IP Address", err)
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	EoE("Error Reading IP Address", err)
	return string(ip)
}
