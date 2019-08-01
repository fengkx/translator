package translator

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/elgs/gojq"
	"github.com/fengkx/translator/config"
	"github.com/imroc/req"
)

type YoudaoTranslator struct {
	name    string
	apihost string
}

type youdaoResp struct {
	DefaultResp
}

func NewYoudaoTranslator(args ...string) Translator {
	var host = config.DefaultYoudaoAPI
	if len(args) > 0 {
		host = args[0]
	}
	return &YoudaoTranslator{
		"youdao",
		host,
	}

}

func (t *YoudaoTranslator) Translate(r Request) (res Respone) {
	var (
		err         error
		definitions map[string]Defintions
	)
	// http://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule
	// i=爱情&from=zh-CHS&to=en&smartresult=dict&client=fanyideskweb
	// &salt=1541752274&sign=99b25d6a0f1dc171047ddf464fa5b87e
	// &doctype=json&version=2.1&keyfrom=fanyi.web&action=FY_BY_REALTIME&typoResult=false

	ms := int64(time.Now().UTC().Unix()) * 1000
	rand.Seed(ms)
	ms = ms + rand.Int63n(10)

	const strStart = "fanyideskweb"
	const strEnd = "ebSeFb%=XZ%T[KZ)c(sy!"
	salt := fmt.Sprintf("%d", ms)

	h := md5.New()
	_, _ = io.WriteString(h, strStart+r.payload+salt+strEnd)
	sign := fmt.Sprintf("%x", h.Sum(nil))

	parm := req.Param{
		"client":      "fanyideskweb",
		"smartresult": "dict",
		"doctype":     "json",
		"version":     "2.1",
		"keyfrom":     "fanyi.web",
		"action":      "FY_BY_REALTIME",
		"typoResult":  "true",
		"from":        r.sl,
		"to":          r.tl,
		"i":           r.payload,
		"salt":        salt,
		"sign":        sign,
	}

	header := req.Header{
		"Cookie":     "OUTFOX_SEARCH_USER_ID=-2022895048@10.168.8.76;",
		"Referer":    "http://fanyi.youdao.com/",
		"User-Agent": "Mozilla/5.0 (Windows NT 6.2; rv:51.0) Gecko/20100101 Firefox/51.0",
	}

	resp, err := req.Post(t.apihost, parm, header)
	if err != nil {
		return &youdaoResp{DefaultResp{err: err, req: r}}
	}

	rawResult, err := resp.ToString()

	jq, err := gojq.NewStringQuery(rawResult)
	if err != nil {
		return &youdaoResp{DefaultResp{err: err, req: r}}
	}

	// result
	result, parseErr := jq.QueryToString("translateResult.[0].[0].tgt")
	source, parseErr := jq.QueryToString("translateResult.[0].[0].src")

	// definiations
	wordMeans, parseErr := jq.QueryToArray("smartResult.entries")

	if parseErr == nil && wordMeans != nil {
		definitions = make(map[string]Defintions)
		posRe := regexp.MustCompile(`(?m)[.]+\s+`)

		for i, wm := range wordMeans {
			wordmean := wm.(string)
			if wordmean == "" {
				continue
			}
			pair := posRe.Split(wordmean, 2)
			if len(pair) > 1 {
				definitions[pair[0]] = Defintions{NewDefintion(strings.TrimSpace(pair[1]))}
			}
			if len(pair) == 1 {
				definitions[fmt.Sprintf("!HIDE!%x", &i)] = Defintions{NewDefintion(strings.TrimSpace(pair[0]))}
			}
		}
	}

	if parseErr != nil {
		err = errors.New(rawResult)
	}

	return &youdaoResp{DefaultResp{
		rawResult:    rawResult,
		src:          source,
		res:          result,
		definitions:  definitions,
		alternatives: nil, // api without any  alternative and translations
		translations: nil,
		req:          r,
		err:          err,
	}}

}

func (t *YoudaoTranslator) Name() string {
	return t.name
}

func (res youdaoResp) print(style bool) {
	if style {
		fmt.Println(ResStyle(res.res))
	} else {
		fmt.Println(res.res)
	}

	if res.translations != nil {
		fmt.Println("--------------------------")

		if style {
			fmt.Println(
				LabelStyle("Translations"))
		} else {
			fmt.Println("Translations")
		}
		for _, v := range res.translations {
			fmt.Printf("\t%s\n", v)
		}
	}
	if res.definitions != nil {
		fmt.Println("--------------------------")
		if style {
			fmt.Println(
				LabelStyle("Definitions"))
		} else {
			fmt.Println("Definitions")
		}
		for k, v := range res.definitions {
			if !strings.HasPrefix(k, "!HIDE!") {
				if style {
					fmt.Printf("[%s]\n", POSStyle(k+"."))
				} else {
					fmt.Printf("[%s]\n", k+".")
				}
			}

			for _, line := range v {
				fmt.Printf("%s", line.string(style))
			}
			fmt.Println()
		}
	}
	if res.alternatives != nil {
		fmt.Println("--------------------------")
		if style {
			fmt.Println(LabelStyle("Alternatives"))
		} else {
			fmt.Println("Alternatives")
		}

		for k, v := range res.alternatives {
			if style {
				fmt.Println(Blod(k))
			} else {
				fmt.Println(k)
			}
			for _, line := range v {
				fmt.Printf("\t%s\n", line)
			}
		}
	}

}

func (res youdaoResp) Print() {
	res.print(true)
}

func (res youdaoResp) RawPrint() {
	res.print(false)
}
