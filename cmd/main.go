package main

import (
	"flag"
	"fmt"
	"github.com/fengkx/translator/translator"
	"os"
	"regexp"
	"strings"
)

func main() {
	var (
		sl string
		tl string
		payload string
	)
	flag.StringVar(&sl, "s", "auto", "source language")
	flag.StringVar(&tl, "t", "", "target language")
	flag.Parse()
	if len(os.Args[1:]) <=0 {
		flag.Usage()
		return
	}
	payload = strings.Join(os.Args[1:], " ")
	if tl == "" {
		re := regexp.MustCompile("[\u4e00-\u9fa5]")
		if re.Match([]byte(payload)) {
			tl = "en"
		} else {
			tl = "zh-CN"
		}
	}

	t := translator.NewGoogleTransaltor()
	result := t.Translate(translator.NewReq(
		payload,
		sl,
		tl))
	if err := result.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	result.Print()
}
