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
	"io"
	"net/http"
 )

 func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("There!")//debugger statement
	destination := r.URL.String()
	response, err := http.Get(destination)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer response.Body.Close()
	_, err = io.Copy(w, response.Body)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
 }

 func main() {
	port := os.Args[1]
	http.HandleFunc("/", handler)
	fmt.Println(port)//making sure port got passed in
	err := (http.ListenAndServe(":"+port, nil))
	if err != nil {
		fmt.Println("Connection error")
	}
 }