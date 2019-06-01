package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/bobesa/go-domain-util/domainutil"
	"github.com/gocolly/colly"
)

// AfterCarw is
func AfterCarw(url string) ResData {
	bb := BaseCraw{}

	methodMap := map[string]interface{}{
		"jd.com":     bb.GetJD2,
		"taobao.com": bb.GetTaoBao,
	}
	v := ""

	domain := domainutil.Domain(url)

	if domain != "" {
		v = domain
	}
	fmt.Println("---------------------")
	fmt.Println(v)
	fmt.Println(methodMap[v])
	if methodMap[v] != nil {
		methodMap[v].(func(string))(url)
	} else {
		bb.ComCarw(url)
	}
	// AppendToFile("jdprive.txt", doc.Text())
	// doc.Find("title").Each(func(i int, selection *goquery.Selection) {
	// 	title = selection.Text()
	// })
	bb.resData.Domain = domain
	return bb.resData
}

// ComCarw is bu'z
func (b *BaseCraw) ComCarw(url string) {
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
