package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/antchfx/htmlquery"
)

func main() {

	header := map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299",
		"Accept-Language": "zh-cn",
	}
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.jdlingyu.com/tuji/page/2", nil)
	if err != nil {
		panic(err)
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 解析响应
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	linkNodes := htmlquery.Find(doc, "//div[@id='post-list']/ul/li")
	//htmlquery.Find  返回值 []*html.Node
	for _, v := range linkNodes {

		htmlStr := htmlquery.FindOne(v, "//h2/a/text()")
		imgName := htmlquery.InnerText(htmlStr)
		imgStr := htmlquery.FindOne(v, "//picture/img")
		imgUrl := htmlquery.SelectAttr(imgStr, "data-src")
		Writelocal(imgUrl, imgName)

	}
}

func Writelocal(url string, imgName string) {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	imgData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(imgName + ".jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	n, err := file.Write(imgData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s保存成功;图片地址%s,共消耗:%d 个字节", imgName, url, n)
	fmt.Println()
}
