package translator

import (
"encoding/json"
"testing"
)

func TestNewGoogleTranslator(t *testing.T) {
	g := NewGoogleTranslator()
	if n:=g.Name(); n != "google" {
		t.Fatalf("GoogleTranslator should name %s but got %s", "google", n)
	}
}

func TestGoogleTranslator_Translate(t *testing.T) {
	g := NewGoogleTranslator()
	tests := [][]string{
		[]string{"Hello", "en", "zh-CN"},
		[]string{"你好", "zh-CN", "en"},
		[]string{"go", "en", "zh-CN"},
	}
	for _, test := range tests {
		resp := g.Translate(NewReq(test[0], test[1], test[2]))
		if resp.Err() != nil {
			t.Fatalf("test fail when translate %s from %s to %s", test[0], test[1], test[2])
		}

		if resp.req.payload != test[0] || resp.req.sl != test[1] || resp.req.tl != test[2] {
			t.Fatalf("test fail req is not eaual to when have passed in")
		}
		if v :=json.Valid([]byte(resp.rawResult)); !v {
			t.Fatalf("test fail raw resp is not json")
		}
	}
}
