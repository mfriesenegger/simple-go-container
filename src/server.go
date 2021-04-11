package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"net"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	hostname, error := os.Hostname()
	if error != nil {
		fmt.Printf("Oops: %v", error)
	}

	addrs, error := net.LookupHost(hostname) 
        if error != nil {
                fmt.Printf("Oops: %v", error)
        }
	fmt.Fprintf(w, "Hello from %s at %s!\n", hostname, addrs)

	out1, error := exec.Command("grep", "PRETTY_NAME", "/etc/os-release").Output()
        if error != nil {
                fmt.Printf("Oops: %v", error)
        }
	fmt.Fprintf(w, "%s", out1)

	out2, error := exec.Command("uname", "-a").Output()
        if error != nil {
                fmt.Printf("Oops: %v", error)
        }
	fmt.Fprintf(w, "%s", out2)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
