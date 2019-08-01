package translator

import (
	"testing"
)

func TestNewBaiduFanyiTranslator(t *testing.T) {
	bdfy := NewBaiduFanyiTranslator()
	if n := bdfy.Name(); n != "bdfanyi" {
		t.Fatalf("BaiduFanyiTranslator should name %s but got %s", "bdfanyi", n)
	}
}

func TestBaiduFanyiTranslator_Translate(t *testing.T) {
	bdfy := NewBaiduFanyiTranslator("", "")
	resp:=bdfy.Translate(NewReq(
		"go",
		"auto",
		"zh",
		))
	if resp.Err() == nil {
		t.Fatal("baidu fanyi require api key")
	}
}