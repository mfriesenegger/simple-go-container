package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"net"
	"regexp"
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

	fileToRead, error := os.Open("/etc/os-release")
	if error != nil {
		log.Fatal(error)
	}
	defer fileToRead.Close()
	scanner := bufio.NewScanner(fileToRead)
	for scanner.Scan() {
		matched, _ := regexp.MatchString("^PRETTY_NAME", scanner.Text())
		if matched {
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


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
