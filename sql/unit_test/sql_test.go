package unit_test

import (
	mysqlconnect "sql/src/mysql_connect"
	"testing"
)

func TestText_extractor(t *testing.T) {

	arr := []string{"url", "any-content"}
	mysqlconnect.Raw_news_insert(arr[0], arr[1])

}
