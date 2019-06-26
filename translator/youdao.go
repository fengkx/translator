package translator

import (
	"crypto/md5"
	"fmt"
	"github.com/elgs/gojq"
	"github.com/imroc/req"
	"io"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type YoudaoTranslator struct {
	name    string
	apihost string
}

func NewYoudaoTranslator(args ...string) Translator {
	var host = "http://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule"
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
	io.WriteString(h, strStart+r.payload+salt+strEnd)
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
	rawResult, err := resp.ToString()
	jq, err := gojq.NewStringQuery(rawResult)

	// result
	result, err := jq.QueryToString("translateResult.[0].[0].tgt")
	source, err := jq.QueryToString("translateResult.[0].[0].src")

	// definiations
	wordMeans, err := jq.QueryToArray("smartResult.entries")

	if err == nil && wordMeans != nil {
		definitions = make(map[string]Defintions)
		posRe := regexp.MustCompile(`(?m)[.]+\s+`)

		for _, wm := range wordMeans {
			wordmean := wm.(string)
			if wordmean == "" {
				continue
			}
			pair := posRe.Split(wordmean, 2)
			definitions[pair[0]] = Defintions{NewDefintion(strings.TrimSpace(pair[1]))}

		}
	}

	if err != nil {
		return Respone{err: err, req: r}
	}

	return Respone{
		rawResult:    rawResult,
		src:          source,
		res:          result,
		definitions:  definitions,
		alternatives: nil, // api without any  alternative and translations
		translations: nil,
		req:          r,
		err:          err,
	}

}

func (t *YoudaoTranslator) Name() string {
	return t.name
}
