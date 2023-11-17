package http

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func readRootCAs(caCertPath string) *x509.CertPool {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// read in the cert file
	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("failed to append %q to RootCAs: %v", "ca.crt", err)

		return rootCAs
	}

	if ok := rootCAs.AppendCertsFromPEM(caCert); !ok {
		log.Println("no certs appended, using system certs only")
	}

	return rootCAs
}

func GetClientPair(crtPath string, keyPath string) (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(crtPath, keyPath)
	if err != nil {
		log.Fatalf("Error opening cert file %s and key %s, Error: %s", crtPath, keyPath, err)
	}

	return &cert, err
}

func JsonPrettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

func JsonPrettyPrint(j any) {
	var buffer bytes.Buffer

	err := JsonPrettyEncode(j, &buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buffer.String())
}

func JsonPrettyPrintV2(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func First[T, U any](val T, _ U) T {
	return val
}

// exists returns whether the given file or directory exists
func Exist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
