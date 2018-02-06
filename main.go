package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	// 引数定義
	t_url := flag.String("u", "", "Please give a URL")
	attr := flag.String("e", "", "Please give a HTML elemet ")
	flag.Parse()

	// ページを取得
	doc, err := goquery.NewDocument(*t_url)
	if err != nil {
		fmt.Print("URL Scraping faild")
	}

	// 記事タイトルだけを抽出
	doc.Find(*attr).Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}
