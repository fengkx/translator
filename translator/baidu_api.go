package translator

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/elgs/gojq"
	"github.com/fengkx/translator/config"
	"github.com/imroc/req"
	"io"
	"math/rand"
	"strconv"
	"time"
)

type BaiduFanyiTranslator struct {
	name    string
	apihost string
	appid string
	appkey string
}

func (t *BaiduFanyiTranslator) Name() string {
	return t.name
}

type BaiduFanyiResp struct {
	req       Request
	rawResult string
	src       string
	res       string
	err error
}

func (res *BaiduFanyiResp) Print() {
	fmt.Println(ResStyle(res.res))
}

func (res *BaiduFanyiResp) RawPrint() {
	fmt.Println(res.res)
}

func (res *BaiduFanyiResp) Err() error {
	return res.err
}


func NewBaiduFanyiTranslator(args ...string) Translator {
	var host = config.DefaultBaiduFanyiAPI
	var (
		appid string
		appkey string
		)
	length := len(args)
	if length == 3 {
		host = args[0]
		appid = args[1]
		appkey = args[2]
	} else if length == 2{
		appid = args[0]
		appkey = args[1]
	}
	return &BaiduFanyiTranslator{
		"bdfanyi",
		host,
		appid,
		appkey,
	}
}

func (t *BaiduFanyiTranslator)Translate(r Request) Respone {
	var err error
	rand.Seed(time.Now().UnixNano())
	salt := strconv.Itoa(rand.Int())
	s := t.appid + r.payload + salt + t.appkey
	m := md5.New()
	io.WriteString(m, s)
	parm := req.Param{
		"from": r.sl,
		"to": r.tl,
		"q": r.payload,
		"appid": t.appid,
		"salt": salt,
		"sign": fmt.Sprintf("%x", m.Sum(nil)),
	}
	resp, err := req.Get(t.apihost, parm)
	if err != nil {
		return &BaiduFanyiResp{err: err, req: r}
	}
	rawResult, err := resp.ToString()
	if err != nil {
		return &BaiduFanyiResp{err: err, req: r}
	}

	jq, err := gojq.NewStringQuery(rawResult)
	if err != nil {
		return &BaiduFanyiResp{err: err, req: r}
	}
	r.sl, err = jq.QueryToString("from")
	r.tl, err = jq.QueryToString("to")
	trans_result, err := jq.QueryToMap("trans_result.[0]")
	if err != nil {
		return &BaiduFanyiResp{err: errors.New(rawResult), req: r}
	}
	src := trans_result["src"].(string)
	res := trans_result["dst"].(string)
	return &BaiduFanyiResp{
		req: r,
		rawResult:rawResult,
		src:src,
		res:res,
	}
}

