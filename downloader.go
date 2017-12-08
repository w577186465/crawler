package crawler

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

type Request struct {
	Method string
}

// 简单下载器
func Download(url string) (*http.Response, bool) {
	resp, err := http.Get(url)

	req, err := http.NewRequest("GET", url, nil) // 发起请求

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println("网页打开失败")
		return nil, false
	}

	return resp, true
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
