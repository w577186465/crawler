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

func Test_json(t *testing.T) {
	json, _ := Download("http://localhost/test/").Json()
	title, _ := json.Get("title").String()
	fmt.Println(title)
}

func Test_jsonp(t *testing.T) {
	url := "http://money.finance.sina.com.cn/d/api/openapi_proxy.php/?__s=[[%22hq%22,%22hs_a%22,%22%22,0,2,40]]&callback=FDC_DC.theTableData"
	json, err := Download(url).Jsonp()
	if err != nil {
		fmt.Println(err)
		return
	}
	count, _ := json.GetIndex(0).Get("day").String()
	fmt.Println(count)
}
