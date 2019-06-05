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

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.DOM.Find("img").Each(func(i int, se *goquery.Selection) {
			fmt.Println(i)

			if i == 0 {
				img, bool := se.Attr("src")
				if bool {
					b.resData.Thumb = img
				}
			}
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
