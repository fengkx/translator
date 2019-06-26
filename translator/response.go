package translator

import "fmt"

type Respone struct {
	req Request
	rawResult string
	src string
	res string
	alternatives map[string][]string
	translations []string
	definitions map[string][]string
	err error
}

func (res *Respone) Err() error {
	return res.err
}

func (res *Respone) Print() {
	fmt.Println(res.res)
	if res.translations != nil {
		fmt.Println("--------------------------")
		fmt.Println("Translations")
		for _,v := range res.translations {
			fmt.Printf("\t%s\n", v)
		}
	}
	if res.definitions != nil {
		fmt.Println("--------------------------")
		fmt.Println("Definitions")
		for k,v := range res.definitions {
			fmt.Println(k)
			for _, line := range v{
				fmt.Printf("\t%s\n", line)
			}
		}
	}
	if res.alternatives != nil {
		fmt.Println("--------------------------")
		fmt.Println("Alternatives")

		for k,v := range res.alternatives {
			fmt.Println(k)
			for _, line := range v{
				fmt.Printf("\t%s\n", line)
			}
		}
	}
}
