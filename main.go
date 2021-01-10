package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/steelx/extractlinks"
)

// this bypasses the certificate

var (
	config = &tls.Config{
		InsecureSkipVerify: true,
	}
	transport = &http.Transport{
		TLSClientConfig: config,
	}
	netClient = &http.Client{
		Transport: transport,
	}
)

func main() {

	arguments := os.Args[1:]
	fmt.Println("args", arguments)
	// baseURLL := "https://www.amazon.com/s?k=hello&ref=nb_sb_noss_2"

	if len(arguments) == 0 {
		fmt.Println("Missing URL, e.g go-webscrapper http://js.org")
		os.Exit(1)
	}

	baseURL := arguments[0]
	fmt.Println("baseURL: ", baseURL)

	crawlURL(baseURL)

}

func crawlURL(href string) {
	fmt.Printf("crawling url -> %+v \n", href)
	response, err := netClient.Get(href)
	checkErr(err)

	defer response.Body.Close()

	links, err := extractlinks.All(response.Body)
	checkErr(err)
	// fmt.Println("LINK: ", link.Href)

	for _, link := range links {
		fmt.Println("LINK: ", link.Href)
		// crawlURL(link.Href)
	}
}

func toFixedURL(href, baseURL string) string {
	uri, err := url.Parse(href)
	// if err != url.Parse(href) {
	// 	if err != nil {
	// 		return " "
	// 	}
	// }

	fmt.Println("URI host", uri.Host)

	base, err := url.Parse(baseURL)
	// if err != url.Parse(href) {
	// 	if err != nil {
	// 		return " "
	// 	}
	// }

	fmt.Println("base host", base.Host)

	return base.String()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
