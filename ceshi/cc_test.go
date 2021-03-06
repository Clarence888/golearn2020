package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestNiHao(t *testing.T) {
	start := time.Now()
	ch := make(chan string)
	for i := 1; i < 10; i++ {
		go fetch("http://129.226.139.199:9191/111", ch) // start a goroutine
	}
	for i := 1; i < 10; i++ {
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
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
