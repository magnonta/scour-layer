// Scour-layer
// Created by: https://github.com/magnonta

package main

import (
	"flag"
	"fmt"
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

func main() {
	flag.Parse()
	var pr = *prt
	var ur = *uri

	// var x string

	// if os.Args[1] == "help" {
	// 	x = help()
	// }

	// fmt.Printf("%v", x)

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
		t, s := sendGet(ur, os.Args[2], pr, port)

		switch port {
		case "80":
			fmt.Printf("pingando http://%s:%s seq=%d time=%s bytes=%d\n", os.Args[2], port, seq, strconv.FormatFloat(t, 'f', 3, 64), s)
		case "443":
			fmt.Printf("pingando https://%s:%s seq=%d time=%s bytes=%d\n", os.Args[2], port, seq, strconv.FormatFloat(t, 'f', 3, 64), s)
		}
		time.Sleep(1 * time.Second)
	}

}
