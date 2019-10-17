package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {

	var url string
	var proxies string
	var ref string
	flag.StringVar(&url, "url", "", "URL to scrape")
	flag.StringVar(&proxies, "proxy", "", "Proxy (optional). Syntax: protocol://ip:port")
	flag.StringVar(&ref, "ref", "", "Referer (optional)")
	flag.Parse()
	c := colly.NewCollector(colly.UserAgent(RandomString()))

	if proxies != "" {
		rp, err := proxy.RoundRobinProxySwitcher(proxies)

		if err != nil {
			log.Fatal(err)
		}

		c.SetProxyFunc(rp)
	}

	c.OnRequest(func(r *colly.Request) {
		if ref != "" {
			r.Headers.Set("Referer", "https://www.displayspecifications.com/en")
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
	})

	c.Visit(url)
}
