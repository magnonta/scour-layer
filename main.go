package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func check(url, port string) (urltime float64, urlsize int) {
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

			url_size := len(body)
			//fmt.Printf("%s", body)
			msec := time.Since(t0)
			url_time := msec.Seconds() * float64(time.Second/time.Millisecond)

			return url_time, url_size
		}
	}
	return
}

func main() {
	fmt.Printf("GoLang httping - PINGING %s\n", os.Args[1])

	port := "80"
	// Do we have port defined ?
	if os.Args[2] != "" {
		port = os.Args[2]
	}

	seq := 0
	// infinite loop
	for {
		seq = seq + 1
		t, s := check(os.Args[1], port)
		fmt.Printf("pingando %s:%s, seq=%d time=%s bytes=%d\n", os.Args[1], port, seq, strconv.FormatFloat(t, 'f', 3, 64), s)
		time.Sleep(1 * time.Second)
	}

}
