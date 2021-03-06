package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/mnorbi/gosandbox/iowordfreq"
)

func readUrls(fileName string) (urls []string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	fs := bufio.NewScanner(f)
	for fs.Scan() {
		urls = append(urls, fs.Text())
	}
	return
}

var start = time.Now()
var iowf struct {
	sync.Mutex
	iowordfreq.IoWordFreq
}
var wordsToKnow = strings.Split("html prize body script css", " ")

func main() {
	fileName := "more_urls.txt"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}
	urls := readUrls(fileName)
	done := make(chan struct{})
	for _, url := range urls {
		go run((&Fetcher{url}).fetch, done)
	}

	for range urls {
		<-done
	}

	fmt.Printf("%.2fs Total\n", time.Since(start).Seconds())
	fmt.Println("==========================")
	for _, w := range wordsToKnow {
		fmt.Printf("%s: %d\n", w, iowf.GetWordCount(w))
	}
}

func run(fn func(), queue chan struct{}) {
	fn()
	queue <- struct{}{}
}

type Fetcher struct {
	Url string
}

func (f *Fetcher) fetch() {
	res, err := http.Get(f.Url)
	if err != nil {
		log.Fatal(err)
		return
	}
	if res.StatusCode != http.StatusOK {
		log.Println("Status eror: ", res)
		return
	}

	iowf.Lock()
	nbytes, err := io.Copy(&iowf, res.Body)
	iowf.Unlock()

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%.2fs  %7d  %s\n", time.Since(start).Seconds(), nbytes, f.Url)
}
