package main

import (
	// "crypto/tls"
	// "crypto/x509"

	// "encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/tsemach/go-load/restapi/restapi"
)

// var host = "localhost"
// var port = 8080
// var caCert = "certs/ca.crt"
// var certOpt = tls.RequireAndVerifyClientCert

// func getTLSConfig(host, caFullPath string, certOpt tls.ClientAuthType) *tls.Config {
// 	var caCert []byte
// 	var err error
// 	var caCertPool *x509.CertPool

// 	if certOpt > tls.RequestClientCert {
// 		caCert, err = os.ReadFile(caFullPath)
// 		if err != nil {
// 			log.Fatal("ERROR opening cert file", caFullPath, ", error ", err)
// 		}
// 		caCertPool = x509.NewCertPool()
// 		caCertPool.AppendCertsFromPEM(caCert)
// 	}

// 	return &tls.Config{
// 		ServerName: host,
// 		ClientAuth: certOpt,
// 		ClientCAs:  caCertPool,
// 		MinVersion: tls.VersionTLS12, // TLS versions below 1.2 are considered insecure - see https://www.rfc-editor.org/rfc/rfc7525.txt for details
// 	}
// }

// func createServer() *http.Server {
// 	g := createGin()
// 	tls := getTLSConfig(host, caCert, tls.ClientAuthType(certOpt))

// 	return &http.Server{
// 		Addr:      fmt.Sprintf(":%v", port),
// 		Handler:   g,
// 		TLSConfig: tls,

// 		ReadTimeout:  20 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 	}
// }

func main() {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)

	certCrt := strings.Join([]string{mydir, "../certs/server.crt"}, "/")
	certKey := strings.Join([]string{mydir, "../certs/server.key"}, "/")

	// create a WaitGroup
	wg := new(sync.WaitGroup)

	// add two goroutines to `wg` WaitGroup
	wg.Add(1)

	// goroutine to launch a server on port 8080
	go func() {
		server := restapi.CreateServer()
		fmt.Println(server.ListenAndServeTLS(certCrt, certKey))
		wg.Done()
	}()

	// wait until WaitGroup is done
	wg.Wait()
}
