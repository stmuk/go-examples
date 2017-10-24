package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {

	r, err := http.Get("http://www.yahoo.com")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	r.Body.Close()
	rd := bytes.NewReader(b)

	doc, err := html.Parse(rd)
	if err != nil {
		log.Fatal(err)
	}

	links := doit(doc)

	fmt.Println("%+v", links)

	for _, v := range links {
		fmt.Printf("%v", v)
	}

}

func doit(n *html.Node) []string {
	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Printf("%s", a.Val)
				links = append(links, a.Val)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			doit(c)
		}
	}
	return links
}
