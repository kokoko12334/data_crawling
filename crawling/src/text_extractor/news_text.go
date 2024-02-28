package text_extractor

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func removeTags(doc *goquery.Selection, tagsToRemove []string) {
	for _, tag := range tagsToRemove {
		doc.Find(tag).Remove()
	}
}

func Extractor(url_ string) string {

	// // 프록시 서버 주소 설정
	// proxyStr := "http://34.64.141.116:3128"
	// proxyURL, err := url.Parse(proxyStr)
	// if err != nil {
	// 	panic(err)
	// }

	// // http.Transport에 프록시 설정
	// transport := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyURL),
	// }

	// // http.Client에 위에서 설정한 Transport 사용
	// client := &http.Client{
	// 	Transport: transport,
	// }

	client := &http.Client{}

	req, err := http.NewRequest("GET", url_, nil)

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
		log.Fatal(err)

	}
	// fmt.Println(doc.Text())
	check := doc.Find("main")
	var doc_content *goquery.Selection
	if check.Text() == "" {
		doc_content = doc.Find("body")
	} else {
		doc_content = check
	}

	// 특정 태그들을 삭제
	removeTags(doc_content, []string{"a", "script", "iframe", "button", "footer", "input", "aside", "form", "img", "em", "small", "i", "li", "ul"})

	// 변경된 HTML 출력
	htmlString := doc_content.Text()
	// fmt.Println(htmlString)
	if err != nil {
		log.Println("Error getting HTML:", err)
		return ""
	}
	result1 := strings.Replace(htmlString, "\t", "", -1)
	result2 := strings.Replace(result1, "\n", "", -1)
	f_result := strings.Replace(result2, " ", "", -1)

	return f_result

}
