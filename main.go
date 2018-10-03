package main

import (
	"fmt"
	"net/http"
	"os"
)

# Environment variables NODE_IP and PORT must be set or passed in

var version = 3
var color = "green"
var node = os.Getenv("NODE_IP")
var intro = fmt.Sprintf(
	"<!DOCTYPE html>"+
		"<html>"+
		"<body>"+

		"<h1 style=\"color:%s;\">Welcome! This is <i>version %d</i> of your application!</h1>"+
		"<h1 style=\"color:%s;\">You are on node %v"+

		"</body>"+
		"</html>", color, version, color, node)

func giveIntro(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, intro)
}

func main() {
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", port),
		Handler: mux,
	}
	mux.HandleFunc("/", giveIntro)
	server.ListenAndServe()
}
