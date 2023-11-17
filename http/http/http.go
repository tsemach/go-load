package http

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"os"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type HTTP[REQ any, RES any] struct {
	client      *http.Client
	req         *http.Request
	reqBody     *REQ
	resBody     *RES
	Url         string
	Path        string
	ContentType string
	rootCAs     *x509.CertPool
	clientPair  *tls.Certificate
}

func NewHTTP[REQ any, RES any](url string) *HTTP[REQ, RES] {
	var http HTTP[REQ, RES]

	http.Url = url
	http.ContentType = "application/json"

	return &http
}

func (h *HTTP[REQ, RES]) SetRootCA(caCertPath string) (*x509.CertPool, error) {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// read in the cert file
	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("failed to append %q to RootCAs: %v", "ca.crt", err)

		return nil, err
	}

	if ok := rootCAs.AppendCertsFromPEM(caCert); !ok {
		log.Println("no certs appended, using system certs only")
	}
	h.rootCAs = rootCAs

	return h.rootCAs, nil
}

func (h *HTTP[REQ, RES]) SetCerts(caCertPath string, certPath string, keyPath string) error {
	_, err := h.SetRootCA(caCertPath)
	if err != nil {
		return err
	}

	_, err = h.SetClientPair(certPath, keyPath)
	if err != nil {
		return err
	}

	return nil
}

func (h *HTTP[REQ, RES]) SetClientPair(certPath string, keyPath string) (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatalf("Error opening cert file %s and key %s, Error: %s", certPath, keyPath, err)
	}
	h.clientPair = &cert

	return &cert, err
}

func (h *HTTP[REQ, RES]) NewRequest(method string, headers *map[string]string, body *REQ) error {
	marshalBody, _ := json.Marshal(*body)
	buffBody := bytes.NewBuffer(marshalBody)
	var err error

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			ServerName:         getDomain(h.Url),
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{*h.clientPair},
			RootCAs:            h.rootCAs,
		},
	}
	h.client = &http.Client{Transport: tr, Timeout: 30 * time.Second}

	h.req, err = http.NewRequest(method, h.Url, buffBody)
	if err != nil {
		return err
	}

	for key, value := range *headers {
		h.req.Header.Add(key, value)
	}

	return nil
}

func (h *HTTP[_, _]) Get(headers *map[string]string) (*http.Response, error) {
	var err error

	if headers == nil {
		headers = &map[string]string{}
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			ServerName:         getDomain(h.Url),
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{*h.clientPair},
			RootCAs:            h.rootCAs,
		},
	}
	h.client = &http.Client{Transport: tr, Timeout: 30 * time.Second}
	h.req, err = http.NewRequest(http.MethodGet, h.Url, nil)

	if err != nil {
		return nil, err
	}

	for key, value := range *headers {
		h.req.Header.Add(key, value)
	}

	if err != nil {
		return nil, err
	}

	resp, err := h.client.Do(h.req)
	if err != nil {
		panic(err)
	}

	return resp, nil
}

func (h *HTTP[REQ, RES]) Post(headers *map[string]string, body *REQ) (*http.Response, error) {
	err := h.NewRequest(http.MethodPost, headers, body)
	if err != nil {
		return nil, err
	}

	resp, err := h.client.Do(h.req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp, nil
}

func getDomain(_url string) string {
	u, _ := url.Parse(_url)
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host)
}

// func (h *HTTP[_, RES]) Send() (*http.Response, *RES) {
// 	resp, err := h.client.Do(h.req)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	var rs RES
// 	err = json.NewDecoder(resp.Body).Decode(&rs)

// 	return resp, &rs
// }

// func (h *HTTP[REQ, _]) PostOld(headers *map[string]string, body *REQ) {
// 	marshalBody, _ := json.Marshal(*body)
// 	buffBody := bytes.NewBuffer(marshalBody)

// 	req, err := http.NewRequest(http.MethodPost, h.Url, buffBody)

// 	if err != nil {
// 		panic(err)
// 	}

// 	for key, value := range *headers {
// 		req.Header.Add(key, value)
// 	}

// 	h.req = req
// }

// func (h *HTTP[REQ, _]) Put(headers *map[string]string, body *REQ) {
// 	marshalBody, _ := json.Marshal(*body)
// 	buffBody := bytes.NewBuffer(marshalBody)

// 	req, err := http.NewRequest(http.MethodPut, h.Url, buffBody)

// 	if err != nil {
// 		panic(err)
// 	}

// 	for key, value := range *headers {
// 		req.Header.Add(key, value)
// 	}

// 	h.req = req
// }

// func (h *HTTP[_, _]) Get(headers *map[string]string) {

// 	req, err := http.NewRequest(http.MethodGet, h.Url, nil)

// 	if err != nil {
// 		panic(err)
// 	}

// 	for key, value := range *headers {
// 		req.Header.Add(key, value)
// 	}

// 	h.req = req
// }

// func GetRootCAs(cacertPath string) *x509.CertPool {
// 	rootCAs, _ := x509.SystemCertPool()
// 	if rootCAs == nil {
// 		rootCAs = x509.NewCertPool()
// 	}

// 	// read in the cert file
// 	caCert, err := os.ReadFile(cacertPath)
// 	if err != nil {
// 		log.Fatalf("failed to append %q to RootCAs: %v", "ca.crt", err)

// 		return rootCAs
// 	}

// 	if ok := rootCAs.AppendCertsFromPEM(caCert); !ok {
// 		log.Println("no certs appended, using system certs only")
// 	}

// 	return rootCAs
// }

// func SendPost[REQ any, RES any](env string, url string, http *HTTP[any, any], body *REQ) (*http.Response, *RES) {

// 	client := GetClient(env)

// 	postBody, _ := json.Marshal(*body)
// 	responseBody := bytes.NewBuffer(postBody)
// 	resp, err := client.Post(fmt.Sprintf("%s%s", url, http.Path), http.ContentType, responseBody)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	var ir RES
// 	err = json.NewDecoder(resp.Body).Decode(&ir)

// 	return resp, &ir
// }

// func Download(uri string, filepath string) {
// 	resp, err := http.Get(uri)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	fileHandle, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer fileHandle.Close()

// 	_, err = io.Copy(fileHandle, resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// }
