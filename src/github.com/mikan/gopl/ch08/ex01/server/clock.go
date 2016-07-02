// Copyright 2016 mikan. All rights reserved.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func handleConn(c net.Conn, location *time.Location) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

var port *int = flag.Int("port", 8000, "tcp port number for listen")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}
	timezone := os.Getenv("TZ")
	if timezone == "" {
		fmt.Println("TZ isn't set.")
		return
	}
	location, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Println("Unknown timezone: " + timezone)
		return
	}
	fmt.Println("TZ: " + location.String())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, location) // handle connections concurrently
	}
}
