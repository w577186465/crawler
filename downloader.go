package crawler

import (
	"fmt"
	// "net"
	"net/http"
	// "io/ioutil"
	// "time"
	"strings"
	"github.com/PuerkitoBio/goquery"
	// "github.com/axgle/mahonia"
)

type Request struct {
	Method string
}

type Response struct {
	Response *http.Response
}

// simple downloader
func Download(url string) (*Response) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil) // 发起请求

	host := gethost(url) // get host

	req.Header.Add("Host", host)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36")

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println("网页打开失败")
		return nil
	}

	response := Response{
		Response: resp,
	}

	return &response
}

// get string url host
func gethost (url string) string {
	a1 := strings.Split(url, "//")[1]
    return strings.Split(a1, "/")[0]
}

func (d Response) Json () {

}

func (d Response) Html () (*goquery.Document, error) {
	resp := d.Response
	defer resp.Body.Close()
	return goquery.NewDocumentFromReader(resp.Body)
}

// func (Request) Download(url string) {

// }

// func Download (url string) {
//   r := *request

//   connTimeout := r.ConnTimeout // 连接超时时间
//   timeout := r.Timeout         // 传输请求时间

//   // 默认时间
//   if connTimeout == 0 {
//     connTimeout = 5
//   }

//   if timeout == 0 {
//     timeout = 10
//   }

//   client := &http.Client{
//     Transport: &http.Transport{
//       Dial: func(netw, addr string) (net.Conn, error) {
//         c, err := net.DialTimeout(netw, addr, time.Second*connTimeout) // 设置建立连接超时
//         if err != nil {
//           return nil, err
//         }
//         c.SetDeadline(time.Now().Add(timeout * time.Second)) // 设置发送接收数据超时
//         return c, nil
//       },
//     },
//   }

//   // form
//   formio := r.Data.Encode()
//   form := strings.NewReader(formio)

//   // 请求类型
//   typeIs := r.Type
//   if typeIs == "" {
//     typeIs = "GET"
//   }

//   req, err := http.NewRequest(typeIs, r.Url, form) // 发起请求

//   // header
//   // req.Header = r.Head
//   req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//   if err != nil {
//     return nil, err
//   }
//   return client.Do(req)
// }
