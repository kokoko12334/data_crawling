package maeil

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Request(from_page int, to_page int, ch chan<- string) {
	// 크롤링할 웹페이지의 URL

	client := &http.Client{}

	for page := from_page; page <= to_page; page++ {

		url := fmt.Sprintf("https://www.hankyung.com/koreamarket/news/all-news?page=%d", page)

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			fmt.Println("Error creating request:", err)
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "Mozilla/5.0")

		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Error sending HTTP request:", err)
			log.Fatal(err)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println("Error parsing document:", err)

		}

		first_body := doc.Find("div.inner_list").First()

		first_body.Find("ul").Each(func(index int, element *goquery.Selection) {
			fmt.Println(element.Text())
			element.Find("li").Each(func(index2 int, element2 *goquery.Selection) {
				url, is := element2.Find("a").Attr("href")

				if is {
					ch <- url + " " + "hk"
				}
			})

		})

	}

}
