package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/u", func(c *gin.Context) {
		url := c.DefaultQuery("url", "")
		title := fetchDataDetail(url)

		c.JSON(200, gin.H{
			"url": title + " || " + url,
		})
	})

	bb := BaseCraw{}
	bb.domain = "jd"

	methodMap := map[string]interface{}{
		"Get": bb.Get,
	}
	v := "Get"
	methodMap[v].(func(string))("zanjs")

	r.Run() // listen and serve on 0.0.0.0:8080
	// 测试功能只抓取10页数据
	// for i := 1; i < 10; i++ {
	// 	url := fmt.Sprintf("https://search.jd.com/Search?keyword=mac&enc=utf-8&wq=mac&page=%d", i)
	// 	fetchData(url)
	// }
}

/**
抓取数据
*/
func fetchDataDetail(url string) string {
	fmt.Println(url)
	client := http.Client{}
	request, err := http.NewRequest("GET", url, strings.NewReader("name=cjb"))
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Linux; U; Android 5.1; zh-cn; m1 metal Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko)Version/4.0 Chrome/37.0.0.0 MQQBrowser/7.6 Mobile Safari/537.36")

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	// 使用NewDocumentFromResponse方式获取获取数据，是应为某些网页会有防止爬取限制，需要设置Header防止被限制
	doc, err := goquery.NewDocumentFromResponse(response)
	title := ""
	// doc.Find("title").Each(func(i int, selection *goquery.Selection) {
	// 	title = selection.Text()
	// 	log.Println(title)
	// })
	doc.Find("#choose-attrs").Find(".item").Each(func(i int, selection *goquery.Selection) {
		log.Println("+++++++++++++++++++++++")
		log.Println(i)
		log.Println(selection.Text())
	})

	return title
}

/**
抓取数据
*/
func fetchData(url string) {
	fmt.Println(url)
	client := http.Client{}
	request, err := http.NewRequest("GET", url, strings.NewReader("name=cjb"))
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Linux; U; Android 5.1; zh-cn; m1 metal Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko)Version/4.0 Chrome/37.0.0.0 MQQBrowser/7.6 Mobile Safari/537.36")

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	// 使用NewDocumentFromResponse方式获取获取数据，是应为某些网页会有防止爬取限制，需要设置Header防止被限制
	doc, err := goquery.NewDocumentFromResponse(response)
	/**
	1：获取ID为J_goodsList 的div节点
	2：获取ul节点
	3：获取li节点列表
	*/
	doc.Find("div[id=\"J_goodsList\"]").Find("ul").Find("li").Each(func(i int, selection *goquery.Selection) {
		// 获取class为p-name p-name-type-2 的div节点，然后获取em子节点的文字内容作为商品标题
		title := selection.Find("div[class=\"p-name p-name-type-2\"]").Find("em").Text()
		// 获取class为p-price的节点，然后获取i标签中的文字作为价格
		price := selection.Find(".p-price").Find("i").Text()
		// 列表中有部分内容是广告内容，不属于标准商品数据，这里排除掉
		if len(title) > 1 {
			// 把获取到的数据追加到jdprive.txt 文件中，格式为  商品名称+三个制表符+价格+换行
			AppendToFile("jdprive.txt", title+"\t\t\t\t"+price+"\n")
		}
	})
}
