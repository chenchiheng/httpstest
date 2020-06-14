package main

import (
    "crypto/tls"
    "crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
    caCertPath = "cert/ca.crt"
    clientCertPath = "cert/client.crt"
    clientKeyPath = "cert/client.key"
    req_url = "https://192.168.3.2:8081"	
)

func main() {
    pool := x509.NewCertPool()
	cacrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
	    fmt.Println("Read file error: ", err)
		return
	}
	pool.AppendCertsFromPEM(cacrt)

	clientCrt, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
	    fmt.Println("LoadX509KeyPair failed: ", err)
		return
	}

	tr := &http.Transport{
	    TLSClientConfig: &tls.Config{
		    RootCAs: pool,
			Certificates: []tls.Certificate{clientCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(req_url)
	if err != nil {
	    fmt.Println("Request error: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
