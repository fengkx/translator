package translator

type Request struct {
	sl      string
	tl      string
	payload string
}

func NewReq(payload, sl, tl string) Request {
	return Request{
		sl,
		tl,
		payload,
	}
}
