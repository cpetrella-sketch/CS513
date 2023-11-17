/*****************************************************************************
 * http_proxy.go
 * Names:
 * NetIds:
 *****************************************************************************/

// TODO: implement an HTTP proxy
package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	destination := r.URL.String()
	fmt.Println(destination) //Making sure the URL was passed in clearly.
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("net/http: use last response")
		},
	} //stop redirects from returning 200
	response, err := client.Get(destination)
	fmt.Println((response.Status)) //Double checking the exact status code
	if err != nil {
		w.WriteHeader(response.StatusCode)
		io.WriteString(w, err.Error())
		return
	}
	defer response.Body.Close()
	w.WriteHeader(response.StatusCode)
	if response.StatusCode == http.StatusOK {
		_, err = io.Copy(w, response.Body)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
	}
}

func main() {
	port := os.Args[1]
	http.HandleFunc("/", handler)
	fmt.Println(port) //making sure port got passed in
	err := (http.ListenAndServe(":"+port, nil))
	if err != nil {
		fmt.Println("Connection error")
	}
}
