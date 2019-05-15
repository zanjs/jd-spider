package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// BaseCraw is
type BaseCraw struct {
	domain string
	title  string
}

// Get is
func (b *BaseCraw) Get(s string) {
	fmt.Println(s)
	fmt.Println(b.domain)
	// return b.domain
}

// GetJD is
func (b *BaseCraw) GetJD(doc goquery.Document) {
	fmt.Println(b)
	// title := ""
	doc.Find("title").Each(func(i int, selection *goquery.Selection) {
		b.title = selection.Text()
	})
	// b.title = title
}

// MethodMapFn is
// func MethodMapFn(b BaseCraw) map[string]interface{} {
// 	return map[string]interface{}{
// 		"get":    b.Get,
// 		"jd.com": b.GetJD,
// 	}
// }
