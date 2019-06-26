package translator

import (
	"github.com/elgs/gojq"
	"github.com/imroc/req"
	"regexp"
)

type CibaTranslator struct {
	name    string
	apihost string
}

func NewCibaTranslator(args ...string) Translator {
	var host = "http://fy.iciba.com/ajax.php"
	if len(args) > 0 {
		host = args[0]
	}
	return &CibaTranslator{
		"ciba",
		host,
	}

}

func (t *CibaTranslator) Translate(r Request) (res Respone) {
	var (
		err         error
		definitions map[string]Defintions
	)
	// http://fy.iciba.com/ajax.php?a=fy&f=zh-CN&t=en&w=%E4%BD%A0%E5%A5%BD
	parm := req.Param{
		"a": "fy",
		"f": r.sl,
		"t": r.tl,
		"w": r.payload,
	}
	resp, err := req.Get(t.apihost, parm)
	rawResult, err := resp.ToString()
	jq, err := gojq.NewStringQuery(rawResult)

	if err != nil {
		return Respone{err: err, req: r}
	}

	// result
	result, parseErr := jq.QueryToString("content.out")

	// definiations
	wordMeans, parseErr := jq.QueryToArray("content.word_mean")

	if parseErr == nil && wordMeans != nil {
		definitions = make(map[string]Defintions)
		posRe := regexp.MustCompile(`(?m)[.]+\s+`)

		for _, wm := range wordMeans {
			wordmean := wm.(string)
			pair := posRe.Split(wordmean, 2)
			definitions[pair[0]] = Defintions{NewDefintion(pair[1])}

		}
	}

	return Respone{
		rawResult:    rawResult,
		src:          r.payload,
		res:          result,
		definitions:  definitions,
		alternatives: nil, // api without any  alternative and translations
		translations: nil,
		req:          r,
		err:          err,
	}

}

func (t *CibaTranslator) Name() string {
	return t.name
}
