package crawler

import (
	"fmt"
	"testing"
)

func Test_download(t *testing.T) {
	res := Download("http://www.zgchaye.cn/news/342/")
	fmt.Println(res)
}

func Test_Html(t *testing.T) {
	html, _ := Download("http://www.zgchaye.cn/baike/344p1/").Html()
	title := html.Find("title").Text()
	fmt.Println(title)
}
