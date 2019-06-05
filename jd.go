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

// MethodMapFn is
// func MethodMapFn(b BaseCraw) map[string]interface{} {
// 	return map[string]interface{}{
// 		"get":    b.Get,
// 		"jd.com": b.GetJD,
// 	}
// }
