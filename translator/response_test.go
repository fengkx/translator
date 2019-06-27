package translator

import (
	"regexp"
	"testing"
)

func TestDefintion_String(t *testing.T) {
	egRe := regexp.MustCompile(`eg.`)

	nosent := NewDefintion("meaning")
	nosentStr := nosent.String()
	if egRe.Match([]byte(nosentStr)){
		t.Fatalf("Definition without sentence should not contain eg. but got %s", nosentStr)
	}

	withsent := NewDefintion("meaning", "sentence")
	withsentStr := withsent.String()
	if !egRe.Match([]byte(withsentStr)) {
		t.Fatalf("Definition without sentence should not contain eg. but got %s", withsentStr)
	}

}
