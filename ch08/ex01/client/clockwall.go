// Copyright 2016 mikan. All rights reserved.

// Netcat1 is a read-only TCP client.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var cities []string

func main() {
	ch := make(chan struct{})
	for i, arg := range os.Args {
		if i == 0 {
			go header()
			continue // command path
		}
		kv := strings.Split(arg, "=") // cIty=host:port
		go func(k, v string) {
			dial(k, v)
			ch <- struct{}{}
		}(kv[0], kv[1])
		cities = append(cities, kv[0])
	}
	for range cities {
		<-ch
	}
}

func dial(city, address string) {
	fmt.Println("dial(" + address + ") open")
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte(city))
	if err != nil {
		log.Fatal(err)
	}
	mustCopy(os.Stdout, conn)
	fmt.Println("dial(" + address + ") close")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func header() {
	for {
		fmt.Println("--------")
		for i, _ := range cities {
			fmt.Printf("%s\n", cities[i])
		}
		time.Sleep(1 * time.Second)
	}
}
