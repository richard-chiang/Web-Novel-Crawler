package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	url := "https://tw.m.ixdzs.com/read/148663/p240.html"
	resp, err := http.Get(url)
	defer resp.Body.Close()

	CheckError("cannot retrieve webpage", err)
	parseHTML(resp.Body)

}

func parseHTML(reader io.ReadCloser) (novel []byte, urls [2]string) {
	z := html.NewTokenizer(reader)
	for {
		switch z.Next() {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			tok := z.Token()
			switch tok.Data {
			case "div":
				for _, a := range tok.Attr {
					if a.Key == "class" && a.Val == "page" {
						urls = GetURL(z)
					}
				}
			}
		}
	}
	return
}

func GetURL(z *html.Tokenizer) (urls [2]string) {
	count := 1
	for i := 0; i < 10; i++ {
		z.Next()
		anchor := z.Token()
		for _, a := range anchor.Attr {
			if a.Key == "href" {
				if count == 1 {
					urls[0] = a.Val
				} else if count == 3 {
					urls[1] = a.Val
				}
				count++
			}
		}
	}

	fmt.Println("urls")
	fmt.Println(urls[0])
	fmt.Println(urls[1])
	return
}

func CheckError(msg string, err error) {
	if err != nil {
		fmt.Println(msg)
		log.Fatal(err.Error())
	}
}
