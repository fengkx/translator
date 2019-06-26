package translator

import (
	"github.com/elgs/gojq"
	"github.com/imroc/req"
	"strings"
)

type GoogleTranslator struct {
	name string
	apihost string
}

func NewGoogleTransaltor(args ...string) GoogleTranslator {
	l := len(args)
	var host string
	if l<1 {
		host = "https://translate.googleapis.com/"
	} else {
		host = args[0]
	}
	if host[len(host)-1] != '/' {
		host = host + "/"
	}
	return GoogleTranslator{
		name: "google",
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
		"sl": r.sl,
		"tl": r.tl,
		"q": r.payload,
	}

	header := req.Header{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:59.0) Gecko/20100101 Firefox/59.0",
		"Accept-Encoding": "gzip, deflate",
		"Accept": "*/*",
		"Connection": "keep-alive",
	}

	resp, err := req.Get(url, param, header)
	if err != nil {
		return Respone{err:err, req:r}
	}
	rawResult, err := resp.ToString()
	if err != nil {
		return Respone{err:err, req:r}
	}
	type respJSON [][][]interface{}
	var s respJSON
	_ = resp.ToJSON(&s)

	jq, err := gojq.NewStringQuery(rawResult)
	result, err := jq.QueryToString("[0].[0].[0]")
	source, err := jq.QueryToString("[0].[0].[1]")
	if err != nil {
		return Respone{err:err, req:r}
	}

	var translations []string
	var alternatives map[string][]string
	var definitions map[string][]string

	if trs, err := jq.QueryToArray("[1].[0].[1]"); err == nil && trs != nil {
		translations = make([]string, len(trs))
		for i, each := range trs {
			translations[i]=each.(string)
		}
	}

	if als, err := jq.QueryToArray("[1].[0].[2]");err == nil && als != nil {
		for _, tr := range als {
			if tr, ok := tr.([]interface{}); ok && len(tr) >=2 {
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

	if defsArr, err := jq.QueryToArray("[12].[0]"); err ==nil && defsArr!=nil {
		strs, err := jq.QueryToArray("[12].[0].[1].[0]")
		if err == nil {
			definitions = make(map[string][]string)
			k := defsArr[0].(string)
			strSlice := make([]string, 2)
			strSlice[0] = strs[0].(string)
			strSlice[1] = strs[2].(string) //ignore strs[1]
			definitions[k]=strSlice
		}



	}
	return Respone{
		req:r,
		alternatives:alternatives,
		translations:translations,
		definitions:definitions,
		rawResult:rawResult,
		res: result,
		src: source,
		err: nil,
	}
}

func (t *GoogleTranslator) Name() string {
	return t.name
}