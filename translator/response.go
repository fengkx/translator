package translator

import (
	"fmt"
	"github.com/fengkx/translator/config"
	"github.com/ttacon/chalk"
)

type Respone struct {
	req          Request
	rawResult    string
	src          string
	res          string
	alternatives map[string][]string
	translations []string
	definitions  map[string]Defintions
	err          error
}

type Defintion struct {
	meaning  string
	sentence string
}

var (
	textColor = config.Cfg.TextColor()
	ResStyle   = textColor.NewStyle().WithTextStyle(chalk.Bold).Style
	LabelStyle = textColor.NewStyle().WithBackground(config.Cfg.LabelColor()).WithTextStyle(chalk.Bold).Style
	POSStyle   = textColor.NewStyle().WithTextStyle(chalk.Bold).WithTextStyle(chalk.Italic).Style
	EgStyle    = textColor.NewStyle().WithTextStyle(chalk.Italic).WithForeground(config.Cfg.EgColor()).Style
	Blod       = textColor.NewStyle().WithTextStyle(chalk.Bold).Style
)

func (d Defintion) String() string {
	if d.sentence != "" {
		return fmt.Sprintf("%s\n\teg: %s\n", d.meaning, EgStyle(d.sentence))
	} else {
		return fmt.Sprintf("%s\n", d.meaning)
	}
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

func (res *Respone) Err() error {
	return res.err
}

func (res *Respone) Print() {
	fmt.Println(ResStyle(res.res))
	if res.translations != nil {
		fmt.Println("--------------------------")

		fmt.Println(
			LabelStyle("Translations"))
		for _, v := range res.translations {
			fmt.Printf("\t%s\n", v)
		}
	}
	if res.definitions != nil {
		fmt.Println("--------------------------")
		fmt.Println(
			LabelStyle("Definitions"))
		for k, v := range res.definitions {
			fmt.Printf("[%s]\n", POSStyle(k+"."))
			for _, line := range v {
				fmt.Printf("%s", line)
			}
			fmt.Println()
		}
	}
	if res.alternatives != nil {
		fmt.Println("--------------------------")
		fmt.Println(LabelStyle("Alternatives"))

		for k, v := range res.alternatives {
			fmt.Println(Blod(k))
			for _, line := range v {
				fmt.Printf("\t%s\n", line)
			}
		}
	}
}
