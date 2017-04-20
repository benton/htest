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
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	host := strings.Split(r.Host, ":")[0]
	msg := fmt.Sprintf("Request for %s%s -> OK", host, r.URL.Path)
	log.Println(msg)
	fmt.Fprintf(w, fmt.Sprintf("%s\n", msg))
}
