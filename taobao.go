package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// TaoBao is
func TaoBao(url string) string {
	title := ""
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("head", func(e *colly.HTMLElement) {
		e.DOM.Find("title").Each(func(i int, selection *goquery.Selection) {
			title = selection.Text()
			fmt.Println(title)
		})

	})
	c.Visit(url)

	return title
}
