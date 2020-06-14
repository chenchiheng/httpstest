package main

import (
    "crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

var req_url = "https://192.168.3.2:8081"
var caCertPath = "cert/ca.crt"

func main() {
    pool := x509.NewCertPool()

	cacrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("Read file error: ", err)
		return
	}

	pool.AppendCertsFromPEM(cacrt)
	tr := &http.Transport{
	    TLSClientConfig: &tls.Config{RootCAs: pool},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(req_url)
	if err != nil {
	    fmt.Println("Get error: ", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}


