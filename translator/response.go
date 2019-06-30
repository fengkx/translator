package translator

import (
	"fmt"
	"github.com/fengkx/translator/config"
	"github.com/ttacon/chalk"
)

var (
	textColor  = config.Cfg.TextColor()
	ResStyle   = textColor.NewStyle().WithTextStyle(chalk.Bold).Style
	LabelStyle = textColor.NewStyle().WithBackground(config.Cfg.LabelColor()).WithTextStyle(chalk.Bold).Style
	POSStyle   = textColor.NewStyle().WithTextStyle(chalk.Bold).WithTextStyle(chalk.Italic).Style
	EgStyle    = textColor.NewStyle().WithTextStyle(chalk.Italic).WithForeground(config.Cfg.EgColor()).Style
	Blod       = textColor.NewStyle().WithTextStyle(chalk.Bold).Style
)

type Respone interface {
	Print()
	RawPrint()
	Err() error
}

type DefaultResp struct {
	req          Request
	rawResult    string
	src          string
	res          string
	alternatives map[string][]string
	translations []string
	definitions  map[string]Defintions
	err          error
}

func (res *DefaultResp) Err() error {
	return res.err
}

func (res *DefaultResp) Req() Request {
	return res.req
}

func (res *DefaultResp) RawResult() string {
	return res.rawResult
}

func (res *DefaultResp) print(style bool) {
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
		fmt.Println(
			LabelStyle("Definitions"))
		for k, v := range res.definitions {
			if style {
				fmt.Printf("[%s]\n", POSStyle(k+"."))
			} else {
				fmt.Printf("[%s]\n", k+".")
			}
			for _, line := range v {
				fmt.Printf("%s", line.string(style))
			}
			fmt.Println()
		}
	}
	if res.alternatives != nil {
		fmt.Println("--------------------------")
		fmt.Println(LabelStyle("Alternatives"))

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

func (res *DefaultResp) Print() {
	res.print(true)
}

func (res *DefaultResp) RawPrint() {
	res.print(false)
}

type Defintion struct {
	meaning  string
	sentence string
}

func (d Defintion) string(style bool) string {
	if d.sentence != "" {
		if style {
			return fmt.Sprintf("%s\n\teg: %s\n", d.meaning, EgStyle(d.sentence))
		} else {
			return fmt.Sprintf("%s\n\teg: %s\n", d.meaning, d.sentence)
		}
	} else {
		return fmt.Sprintf("%s\n", d.meaning)
	}
}

func (d Defintion) StyleString() string {
	return d.string(true)
}

func (d Defintion) RawString() string {
	return d.string(false)
}

type Defintions []Defintion

func NewDefintion(args ...string) Defintion {
	l := len(args)
	if l > 1 {
		return Defintion{
			args[0],
			args[1],
		}
	} else {
		return Defintion{
			args[0],
			"",
		}
	}
}
