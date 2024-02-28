package entity

type raw_news_info struct {
	id      uint64
	url     string
	content string
}

func Get_raw_news_info_fileds() []string {
	return []string{"id", "url", "content"}

}
