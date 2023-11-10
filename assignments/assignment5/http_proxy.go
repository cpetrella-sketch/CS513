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
    fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
 }

 func main() {
	port := os.Args[1]
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))//Note: change port number to be not hard coded.
 }