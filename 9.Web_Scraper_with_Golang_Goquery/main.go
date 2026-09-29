package main

import (
	"fmt"
	"log"
	"net/http"

	// Import the goquery package
	"github.com/PuerkitoBio/goquery"
)

func main() {
	blogTitles, err := GetLatestBlogTitles("https://www.scrapingbee.com/blog/")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Blog Titles: ")
	fmt.Printf("%s\n", blogTitles)
}
func GetLatestBlogTitles(url string) (string, error) {

	resp, err := http.Get(url)

	if err != nil {
		return " ", err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return " ", err
	}
	titles := ""
	doc.Find("h4").Each(func(i int, s *goquery.Selection) {
		titles += "-" + s.Text() + "\n"
	})
	return titles, nil
}
