package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// GetTaoBao is
func (b *BaseCraw) GetTaoBao(url string) {
	title := ""
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("head", func(e *colly.HTMLElement) {
		e.DOM.Find("title").Each(func(i int, selection *goquery.Selection) {
			title = selection.Text()
			// b.title = title + "测试"
			b.resData.Title = title + "测试"
			fmt.Println(title)
		})

	})
	c.Visit(url)
}

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
