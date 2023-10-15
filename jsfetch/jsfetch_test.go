package jsfetch

import (
	"fmt"
	"net/http"
	"testing"

	"d1y.io/jslike/jsfetch/multipart"
)

func TestFetch(t *testing.T) {
	form := multipart.NewFormData()
	form.Append("name", "你好世界")
	resp, err := Fetch("https://baidu.com", Options{
		Method: http.MethodPost,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body: form.Body(),
	})
	if err != nil {
		t.FailNow()
	}
	text, err := resp.Text()
	if err != nil {
		t.FailNow()
	}
	fmt.Println(text)
}
