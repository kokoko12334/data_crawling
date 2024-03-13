package main

import (
	"crawling/src/naver"
	"crawling/src/text_extractor"
	"log"
	"math/rand"
	"os"
	"runtime"
	mysql "sql/src/mysql_connect"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	gopath := os.Getenv("GOPATH")
	log_path := gopath + "/exec/logfile.log"
	log_file, err := os.OpenFile(log_path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal("파일을 만들 수 없습니다: ", err)
	}

	defer log_file.Close()

	log.SetOutput(log_file)

	ch_url := make(chan string, 100)
	ch_news := make(chan []string, 100)

	runtime.GOMAXPROCS(3)

	var wait sync.WaitGroup

	wait.Add(3)

	go func() {
		defer wait.Done()

		current_time := time.Now()
		yesterday := current_time.AddDate(0, 0, -1)
		date_string := yesterday.Format("20060102")
		date_int, err := strconv.Atoi(date_string)
		if err != nil {
			log.Println("날짜를 정수로 변환하는 데 문제가 발생했습니다:", err)
			return
		}
		// maeil.Request(1, 1, ch_url)
		// hankyung.Request(1, 1, ch_url)
		log.Printf("뉴스날짜: %d", date_int)
		naver.Request(1, 20, date_int, 401, ch_url)
		naver.Request(1, 20, date_int, 402, ch_url)
		close(ch_url)

	}()

	go func() {
		defer wait.Done()
		for url := range ch_url {
			url_split := strings.Split(url, " ")
			f_url, source := url_split[0], url_split[1]
			log.Println(f_url, source)
			ch_news <- []string{f_url, text_extractor.Extractor(f_url), source}

			rand.New(rand.NewSource(time.Now().UnixNano()))
			// 0.0에서 1.0 사이의 랜덤한 소수 생성
			randomFloat := 0.1 + rand.Float64()*0.9
			// 0.0에서 1.0을 소수점 한 자리까지로 변환
			randomFloatWithOneDecimal := float64(int(randomFloat*10)) / 10
			// Sleep에 사용할 시간 계산 (초 단위)
			sleepTime := time.Duration(randomFloatWithOneDecimal * float64(time.Second))
			log.Printf("Sleeping for %.1f seconds\n", randomFloatWithOneDecimal)
			time.Sleep(sleepTime)

		}

		close(ch_news)
	}()

	go func() {
		defer wait.Done()

		db := mysql.Mysql_connect()
		defer db.Close()

		for url_news := range ch_news {
			mysql.Raw_news_insert(db, url_news[0], url_news[1], url_news[2])

		}

	}()

	wait.Wait()

}
