package translator

import (
	"github.com/elgs/gojq"
	"github.com/imroc/req"
	"strings"

	"github.com/fengkx/translator/config"
)

type GoogleTranslator struct {
	name    string
	apihost string
}

type googleResp struct {
	DefaultResp
}

func NewGoogleTranslator(args ...string) Translator {
	l := len(args)
	var host string
	if l < 1 {
		host = config.DefaultGoogleAPI
	} else {
		host = args[0]
	}
	if host[len(host)-1] != '/' {
		host = host + "/"
	}
	return &GoogleTranslator{
		name:    "google",
		apihost: host,
	}
}

func (t *GoogleTranslator) Translate(r Request) (res Respone) {
	url := t.apihost + "translate_a/single"
	dts := []string{
		"bd",
		"ex",
		"ld",
		"md",
		"qca",
		"rw",
		"rm",
		"ss",
		"t",
	}

	for i, t := range dts {
		dts[i] = "dt=" + t
	}
	dt := strings.Join(dts, "&")
	url = url + "?" + dt
	param := req.Param{
		"client": "gtx",
		"sl":     r.sl,
		"tl":     r.tl,
		"q":      r.payload,
	}

	header := req.Header{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:59.0) Gecko/20100101 Firefox/59.0",
		"Accept-Encoding": "gzip, deflate",
		"Accept":          "*/*",
		"Connection":      "keep-alive",
	}

	resp, err := req.Get(url, param, header)
	if err != nil {
		return &googleResp{DefaultResp{err: err, req: r}}
	}
	rawResult, err := resp.ToString()
	if err != nil {
		return &DefaultResp{err: err, req: r}
	}
	type respJSON [][][]interface{}
	var s respJSON
	_ = resp.ToJSON(&s)

	jq, err := gojq.NewStringQuery(rawResult)
	if err != nil {
		return &googleResp{DefaultResp{err: err, req: r}}
	}
	result, _ := jq.QueryToString("[0].[0].[0]")
	source, _ := jq.QueryToString("[0].[0].[1]")

	var translations []string
	var alternatives map[string][]string
	var definitions map[string]Defintions

	if trs, parseErr := jq.QueryToArray("[1].[0].[2]"); parseErr == nil && trs != nil {
		translations = make([]string, 0, len(trs))
		for _, each := range trs {
			eachSlice := each.([]interface{})
			rawWords := eachSlice[1].([]interface{})
			words := make([]string, len(rawWords))
			for i, word := range rawWords {
				words[i] = word.(string)
			}

			translations = append(translations, eachSlice[0].(string)+": "+strings.Join(words, " "))
		}
	}

	if als, parseErr := jq.QueryToArray("[1].[0].[2]"); parseErr == nil && als != nil {
		for _, tr := range als {
			if tr, ok := tr.([]interface{}); ok && len(tr) >= 2 {
				alternatives = make(map[string][]string)
				k := tr[0].(string)
				vs := tr[1].([]interface{})

				v := make([]string, len(vs))
				for i, s := range vs {
					v[i] = s.(string)
				}
				alternatives[k] = v
			}
		}
	}

	if pos, parseErr := jq.QueryToArray("[12]"); parseErr == nil && pos != nil {
		definitions = make(map[string]Defintions)
		for _, p := range pos {
			if defsArr, ok := p.([]interface{}); ok && defsArr != nil {
				strs := defsArr[1].([]interface{})
				defItems := make(Defintions, len(strs))
				k := defsArr[0].(string)
				for i, d := range strs {
					if df, ok := d.([]interface{}); ok {
						meaning := df[0].(string)
						if len(df) > 2 {
							sentence := df[2].(string)
							defItems[i] = NewDefintion(meaning, sentence)
						} else {
							defItems[i] = NewDefintion(meaning)
						}

					}
					definitions[k] = defItems
				}

			}
		}
	}
	return &googleResp{DefaultResp{
		req:          r,
		alternatives: alternatives,
		translations: translations,
		definitions:  definitions,
		rawResult:    rawResult,
		res:          result,
		src:          source,
		err:          nil,
	}}
}

func (t *GoogleTranslator) Name() string {
	return t.name
}
