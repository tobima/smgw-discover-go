package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/tobima/smgw-discover-go/smgw"
)

func main() {
	// Discover SMGW
	host, err := smgw.Discover()
	if err != nil {
		fmt.Println("Error discovering SMGW:", err)
		return
	}
	fmt.Println("Using host:", host)

	// Create HTTP client with custom transport
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		ForceAttemptHTTP2: false,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://" + host)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
		os.Exit(1)
	}
	fmt.Println("Response Status:", resp.Status)
}
