// Copyright 2015-2016 mikan. All rights reserved.

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nBytes, url)
}

// >bin\ex11.exe https://www.google.com https://www.facebook.com https://www.youtube.com https://www.baidu.com https://www.yahoo.com https://www.amazon.com https://www.wikipedia.org https://www.qq.com
// Get https://www.baidu.com: dial tcp: lookup www.baidu.com: getaddrinfow: This is usually a temporary error during hostname resolution and means that the local server did not receive a response from an authoritative server.
// 0.20s   19354 https://www.google.com
// 0.29s   71955 https://www.facebook.com
// 0.76s  428143 https://www.youtube.com
// 0.77s   49491 https://www.wikipedia.org
// 1.15s  401413 https://www.yahoo.com
// 2.12s  396802 https://www.amazon.com
// Get https://www.qq.com: dial tcp 103.7.30.123:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
// 21.04s elapsed
