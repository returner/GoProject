package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the Html page.
	dorosiwaUrl := "https://m.dorosiwa.co.kr/product/detail.html?product_no=5181&cate_no=260&display_group=1"
	//dorosiwaUrl := "https://m.dorosiwa.co.kr"
	res, err := http.Get(dorosiwaUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("ststus code error: %d %s", res.StatusCode, res.Status)
	}
	//https://cdn.snapfit.co.kr/review/contents/stores/dorosiwa/20210305/medium_d93d78fc8f84f231a58461dea3103127.jpg
	//Load the HTML document
	//fmt.Printf("%v", res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%v", &doc)
	// Find the review items
	wrapper := doc.Find("img")
	//items := wrapper.Find("a")
	wrapper.Each(func(idx int, sel *goquery.Selection) {
		src, ok := sel.Attr("src")
		if ok {
			if strings.Contains(src, "https://cdn.snapfit.co.kr") {
				fmt.Println(src)
			}
		}
	})

	// doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
	// 	//For each item found, get the title
	// 	title := s.Find("a").Text()
	// 	fmt.Printf("Review %d: %s\n", i, title)
	// })
}

func main() {
	ExampleScrape()
}
