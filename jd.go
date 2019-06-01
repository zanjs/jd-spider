package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// BaseCraw is
type BaseCraw struct {
	domain  string
	resData ResData
}

// GetJD2 is
func (b *BaseCraw) GetJD2(url string) {
	title := ""
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("head", func(e *colly.HTMLElement) {
		e.DOM.Find("title").Each(func(i int, selection *goquery.Selection) {
			title = selection.Text()
			b.resData.URL = url
			b.resData.Title = title + "测试"
			fmt.Println(title)
		})

	})
	c.Visit(url)
}

// MethodMapFn is
// func MethodMapFn(b BaseCraw) map[string]interface{} {
// 	return map[string]interface{}{
// 		"get":    b.Get,
// 		"jd.com": b.GetJD,
// 	}
// }
