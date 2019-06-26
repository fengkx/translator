package translator

const (
	DefaultUA = "Mozilla/5.0 translator"
)

type Translator interface {
	Name() string
	Translate(r Request) (res Respone)
}
