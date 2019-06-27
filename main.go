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
	// register translator
	type translatorConstructor func(args ...string) translator.Translator

	var constructors = map[string]translatorConstructor{
		"google": translator.NewGoogleTranslator,
		"ciba":   translator.NewCibaTranslator,
		"youdao": translator.NewYoudaoTranslator,
	}

	var (
		sl      string
		tl      string
		payload string
		engine  string
	)

	flag.StringVar(&sl, "s", "auto", "source language")
	flag.StringVar(&engine, "e", "google", "engine")
	flag.StringVar(&tl, "t", "", "target language")
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		engines := make([]string, 0, len(constructors))
		for k := range constructors {
			engines = append(engines, k)
		}
		fmt.Fprintf(flag.CommandLine.Output(), "Supported engines: %s\n", strings.Join(engines, ", "))

	}
	args := flag.Args()
	if len(args) <= 0 {
		flag.Usage()
		return
	}
	payload = strings.Join(args, " ")
	if tl == "" {
		re := regexp.MustCompile("[\u4e00-\u9fa5]")
		if re.Match([]byte(payload)) {
			tl = "en"
		} else {
			tl = "zh-CN"
		}
	}

	c, ok := constructors[engine]
	if !ok {
		flag.Usage()
		return
	}
	t := c()
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
