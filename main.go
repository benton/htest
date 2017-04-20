package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	port := flag.Int("port", 80, "listening port")
	flag.Parse()
	fmt.Printf("Starting on port %d:...\n", *port)
	http.HandleFunc("/", brokenHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	host := strings.Split(r.Host, ":")[0]
	msg := fmt.Sprintf("Request for %s%s -> OK\n", host, r.URL.Path)
	log.Print(msg)
	fmt.Fprintf(w, msg)
}

func brokenHandler(w http.ResponseWriter, r *http.Request) {
	host := strings.Split(r.Host, ":")[0]
	msg := fmt.Sprintf("Request for %s%s -> NOT OK! ☄\n", host, r.URL.Path)
	log.Print(msg)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(msg))
}
