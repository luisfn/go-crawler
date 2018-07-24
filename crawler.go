package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/antchfx/htmlquery"
	"github.com/luisfn/crawler/crawlers"
)


func main() {

	results := make(chan interface{}, 100)

	links := getLinks("links.txt")

	for _, url := range links {
		go worker(results, url)
	}

	for range links {
		res := <-results
		fmt.Println(res)
	}
}

func worker(results chan interface{}, url string) {
	amazon := crawlers.Amazon{}
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		panic(err)
	}

	results <- amazon.Scrap(doc).Json()
}

func getLinks(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var links []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		links = append(links, scanner.Text())
	}

	return links
}