package crawler

import (
	"fmt"
	"testing"
)

func Test_download(t *testing.T) {
	Download("http://www.zgchaye.cn/news/342/")
}

func Test_Html(t *testing.T) {
	html, _ := Download("http://www.zgchaye.cn/news/342/").Html()
	title := html.Find("title").Text()
	fmt.Println(title)
}