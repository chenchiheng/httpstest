package main

import (
    "fmt"
    "net/http"
)
var listen_ip = "192.168.3.2"
var listen_port = "8081"

func handler(w http.ResponseWriter, r *http.Request){
    fmt.Println("Get request from ", r.RemoteAddr)
    fmt.Fprintf(w, "This is a http/https tester service in golang!")
}

func main(){
    http.HandleFunc("/", handler)
    fmt.Printf("Starting http server on %s:%s...............\n", listen_ip, listen_port)
    http.ListenAndServeTLS(listen_ip+":"+listen_port, "cert/server.crt", "cert/server.key", nil)
}
