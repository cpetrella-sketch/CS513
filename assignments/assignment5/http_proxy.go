/*****************************************************************************
 * http_proxy.go                                                                 
 * Names: 
 * NetIds:
 *****************************************************************************/

 // TODO: implement an HTTP proxy
 package main

 import (
	"fmt"
	"os"
	"net/http"
 )

 func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
 }

 func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
 }