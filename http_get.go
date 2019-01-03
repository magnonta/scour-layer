// Scour-layer
// Created by: https://github.com/magnonta

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// sendGet send request to the uri and return data for analyze
func sendGet(uri, url, prt, port string) (urltime float64, urlsize int) {

	t0 := time.Now()
	client := &http.Client{}

	// default case: HTTP request
	domain := "http://" + url + ":" + port
	switch port {
	case "80":
		domain = "http://" + url
	case "443":
		domain = "https://" + url
	}

	req, err := http.NewRequest("GET", domain, nil)
	if err != nil {
		// handle error
		log.Fatalf("|----- Não foi possível conectar %s", url)
	} else {
		req.Proto = "HTTP/1.1"
		req.ProtoMinor = 0
		req.Header.Set("User-Agent", "GoLang httping v0.1")

		resp, err := client.Do(req)
		if err != nil {
			// handle error
			log.Fatalf("|----- Não foi possível conectar %s\n", url)
		} else {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			urlSize := len(body)
			msec := time.Since(t0)
			urlTime := msec.Seconds() * float64(time.Second/time.Millisecond)

			return urlTime, urlSize
		}
	}
	return
}
