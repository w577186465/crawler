package crawler

import (
	"fmt"
	"testing"
)

func Test_Download(t *testing.T) {
	_, ok := Download("http://localhost/test/index.php")
	if ok {
		fmt.Println("success")
	}
}
