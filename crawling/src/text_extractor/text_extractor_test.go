package text_extractor

import (
	"testing"
)

func TestText_extractor(t *testing.T) {

	result := Extractor("https://n.news.naver.com/mnews/article/015/0004913993")
	if result != "" {
		t.Logf("result:%s", result)
	} else {
		t.Error("errrrr")
	}

}
