package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
)

// CLI Commands to use
var uri = flag.String("uri", "", "Endereço que deseja enviar requisições")
var prt = flag.String("port", "", "80 para HTTP ou 443 para HTTPS")
var hlp = flag.String("help", "", "")

func init() {
	flag.StringVar(uri, "u", "", "")
	flag.StringVar(prt, "p", "", "")
	flag.StringVar(hlp, "h", "", "")
}

func check(uri, url, prt, port string) (urltime float64, urlsize int) {

	t0 := time.Now()
	client := &http.Client{}

	if port == "" {
		port = "80"
	}
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

func main() {
	flag.Parse()
	var pr = *prt
	var ur = *uri

	fmt.Printf("GoLang httping - PINGING %s\n", os.Args[2])

	// Define spinner
	s := spinner.New(spinner.CharSets[23], 100*time.Millisecond)
	s.Color("white")
	s.Start()
	time.Sleep(1 * time.Second)

	port := "80"
	// Do we have port defined ?
	if len(os.Args) < 4 {
		port = "80"
	} else {
		port = os.Args[4]
	}

	// infinite loop
	seq := 0
	for {
		seq = seq + 1
		t, s := check(ur, os.Args[2], pr, port)

		switch port {
		case "80":
			fmt.Printf("pingando http://%s:%s seq=%d time=%s bytes=%d\n", os.Args[2], port, seq, strconv.FormatFloat(t, 'f', 3, 64), s)
		case "443":
			fmt.Printf("pingando https://%s:%s seq=%d time=%s bytes=%d\n", os.Args[2], port, seq, strconv.FormatFloat(t, 'f', 3, 64), s)
		}
		time.Sleep(1 * time.Second)
	}

}
