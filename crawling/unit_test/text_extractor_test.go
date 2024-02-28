package unit_test

import (
	"crawling/src/text_extractor"
	"testing"
)

func TestText_extractor(t *testing.T) {

	result := text_extractor.Extractor("https://n.news.naver.com/mnews/article/015/0004913993")
	if result != "" {
		t.Logf("result:%s", result)
	} else {
		t.Error("errrrr")
	}

}
