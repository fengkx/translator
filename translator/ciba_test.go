package translator

import (
	"encoding/json"
	"testing"
)

func TestNewCibaTranslator(t *testing.T) {
	ciba := NewCibaTranslator()
	if n := ciba.Name(); n != "ciba" {
		t.Fatalf("CibaTranslator should name %s but got %s", "ciba", n)
	}
}

func TestCibaTranslator_Translate(t *testing.T) {
	ciba := NewCibaTranslator()
	tests := [][]string{
		{"Hello", "en", "zh-CN"},
		{"你好", "zh-CN", "en"},
		{"go", "en", "zh-CN"},
	}
	for _, test := range tests {
		r := ciba.Translate(NewReq(test[0], test[1], test[2]))
		resp := r.(interface {
			Err() error
			Req() Request
			RawResult() string
		})
		if resp.Err() != nil {
			t.Fatalf("test fail when translate %s from %s to %s", test[0], test[1], test[2])
		}

		if resp.Req().payload != test[0] || resp.Req().sl != test[1] || resp.Req().tl != test[2] {
			t.Fatalf("test fail req is not eaual to when have passed in")
		}
		if v := json.Valid([]byte(resp.RawResult())); !v {
			t.Fatalf("test fail raw resp is not json")
		}
	}
}
