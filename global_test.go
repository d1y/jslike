package jslike

import (
	"testing"
)

func TestGlobalURL(t *testing.T) {
	var url = `https://爱.love/#dev?query=patch&name=你好世界`
	var result = EncodeURI(url)
	var url2, _ = DecodeURI(result)
	if url != url2 {
		t.FailNow()
	}
}

func TestGlobalBase64(t *testing.T) {
	var raw = "你好"
	var b64 = Btoa(raw)
	var b641, _ = Atob(b64)
	if raw != b641 {
		t.FailNow()
	}
}
