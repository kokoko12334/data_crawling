package main

import (
	"crawling/src/naver"
	"crawling/src/text_extractor"
	"fmt"
	"math/rand"
	"runtime"
	mysqlconnect "sql/src/mysql_connect"
	"strings"
	"sync"
	"time"
)

func main() {

	ch_url := make(chan string, 100)
	ch_news := make(chan []string, 100)

	runtime.GOMAXPROCS(3)

	var wait sync.WaitGroup

	wait.Add(3)

	go func() {
		defer wait.Done()
		// maeil.Request(1, 1, ch_url)
		// hankyung.Request(1, 1, ch_url)
		naver.Request(1, 10, 20240227, ch_url)
		close(ch_url)

	}()

	go func() {
		defer wait.Done()
		for url := range ch_url {
			url_split := strings.Split(url, " ")
			f_url, source := url_split[0], url_split[1]
			fmt.Println(f_url, source)
			ch_news <- []string{f_url, text_extractor.Extractor(f_url), source}

			rand.New(rand.NewSource(time.Now().UnixNano()))
			// 0.0에서 1.0 사이의 랜덤한 소수 생성
			randomFloat := 0.1 + rand.Float64()*0.9
			// 0.0에서 1.0을 소수점 한 자리까지로 변환
			randomFloatWithOneDecimal := float64(int(randomFloat*10)) / 10
			// Sleep에 사용할 시간 계산 (초 단위)
			sleepTime := time.Duration(randomFloatWithOneDecimal * float64(time.Second))
			fmt.Printf("Sleeping for %.1f seconds\n", randomFloatWithOneDecimal)
			time.Sleep(sleepTime)

		}

		close(ch_news)
	}()

	go func() {
		defer wait.Done()

		for url_news := range ch_news {
			mysqlconnect.Raw_news_insert(url_news[0], url_news[1], url_news[2])

		}

	}()

	wait.Wait()
}
