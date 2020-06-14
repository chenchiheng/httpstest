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
    serverCertPath = "cert/server.crt"
    serverKeyPath = "cert/server.key"
	listen_ip = "192.168.3.2"
	listen_port = "8081"
	listen_url = listen_ip + ":" + listen_port
)


type httpshandler struct {
}

func (handler *httpshandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a test https server in go!")
}


func main() {
	pool := x509.NewCertPool()
	cacrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
	    fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(cacrt)

	server := &http.Server{
	    Addr: listen_url,
		Handler: &httpshandler{},
		TLSConfig: &tls.Config{
		    ClientCAs: pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	err = server.ListenAndServeTLS(serverCertPath, serverKeyPath)
	if err != nil {
	    fmt.Println("ListenAndServerTLS err:", err)
	}

}
