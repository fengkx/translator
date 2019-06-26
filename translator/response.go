package translator

import "fmt"

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

func (d Defintion) String() string {
	if d.sentence != "" {
		return fmt.Sprintf("%s\n\teg:%s\n", d.meaning, d.sentence)
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
	fmt.Println(res.res)
	if res.translations != nil {
		fmt.Println("--------------------------")
		fmt.Println("Translations")
		for _, v := range res.translations {
			fmt.Printf("\t%s\n", v)
		}
	}
	if res.definitions != nil {
		fmt.Println("--------------------------")
		fmt.Println("Definitions")
		for k, v := range res.definitions {
			fmt.Printf("[%s]\n", k)
			for _, line := range v {
				fmt.Printf("%s", line)
			}
			fmt.Println()
		}
	}
	if res.alternatives != nil {
		fmt.Println("--------------------------")
		fmt.Println("Alternatives")

		for k, v := range res.alternatives {
			fmt.Println(k)
			for _, line := range v {
				fmt.Printf("\t%s\n", line)
			}
		}
	}
}
