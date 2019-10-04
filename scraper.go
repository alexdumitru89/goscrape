package main

import (
	"log"
	"flag"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"fmt"
)

func main() {

	var url string
	var proxies string
	flag.StringVar(&url, "url", "", "URL to scrape")
	flag.StringVar(&proxies, "proxy", "", "Proxy (optional). Syntax: protocol://ip:port")
	flag.Parse()
	c := colly.NewCollector()

	rp, err := proxy.RoundRobinProxySwitcher(proxies) 
	if err != nil {
		log.Fatal(err)
	}
 
	c.SetProxyFunc(rp)

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response){
		fmt.Println(string(r.Body))
	})

	c.Visit(url)
}
