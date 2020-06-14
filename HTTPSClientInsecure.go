package main

import (
    "fmt"
    "net/http"
    "crypto/tls"
    "io/ioutil"
)

var req_url = "https://192.168.3.2:8081"


func main() {
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    resp, err := client.Get(req_url)

    if err != nil {
        fmt.Println("error: ", err)
	return
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
