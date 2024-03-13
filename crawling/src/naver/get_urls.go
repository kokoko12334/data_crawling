package naver

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Request(from_page int, to_page int, date int, section_id3 int, ch chan<- string) {

	client := &http.Client{}

	for page := from_page; page <= to_page; page++ {

		url := fmt.Sprintf("https://finance.naver.com/news/news_list.naver?mode=LSS3D&section_id=101&section_id2=258&section_id3=%d&date=%d&page=%d", section_id3, date, page)

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			log.Println("Error creating request:", err)
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "Mozilla/5.0")
		resp, err := client.Do(req)

		if err != nil {
			log.Println("Error sending HTTP request:", err)
			log.Fatal(err)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Println("Error parsing document:", err)

		}

		first_doc := doc.Find("ul.realtimeNewsList")

		if text := strings.TrimSpace(first_doc.Text()); text == "" {
			log.Println("end")
			break
		}

		first_doc.Find("dt.thumb").Each(func(index int, element *goquery.Selection) {
			sub_url, is := element.Find("a").Attr("href")
			parts1 := strings.Split(sub_url, "?")
			sub_url2 := strings.Split(parts1[1], "&")
			office_id := strings.Split(sub_url2[1], "=")[1]
			article_id := strings.Split(sub_url2[0], "=")[1]
			final_url := fmt.Sprintf("https://n.news.naver.com/mnews/article/%s/%s", office_id, article_id)
			if len(final_url) == 0 {
				log.Println("end")
			}
			if is {
				ch <- final_url + " " + "naver"
			}

		})

	}

}
