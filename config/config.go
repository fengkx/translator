package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/ttacon/chalk"
	"io"
	"log"
	"os"
	"path"
)

const (
	DefaultGoogleAPI = "https://translate.googleapis.com/"
	DefaultYoudaoAPI = "http://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule"
	DefaultCibaAPI   = "http://fy.iciba.com/ajax.php"
)

var ColorMap = map[string]chalk.Color{
	"black":      chalk.Black,
	"red":        chalk.Red,
	"green":      chalk.Green,
	"yellow":     chalk.Yellow,
	"blue":       chalk.Blue,
	"magenta":    chalk.Magenta,
	"cyan":       chalk.Cyan,
	"white":      chalk.White,
	"resetColor": chalk.ResetColor,
}

const (
	DefaultLabelColor = "green"
	DefaultTextColor  = "white"
	DefaultEgColor    = "yellow"
)

const filename = "go-translator.ini"

var inipath = path.Join(ConfigPath(), filename)

type Config struct {
	*ini.File
}

var Cfg Config

func (c Config) TextColor() chalk.Color {
	return ColorMap[c.Section("output").Key("TextColor").Value()]
}

func (c Config) EgColor() chalk.Color {
	return ColorMap[c.Section("output").Key("EgColor").Value()]
}

func (c Config) LabelColor() chalk.Color {
	return ColorMap[c.Section("output").Key("LabelColor").Value()]
}

const tpl = `# Translator configuration
[google]
HOST=%s

[ciba]
HOST=%s

[youdao]
HOST=%s

[output]
# only support black red green yellow blue magenta cyan white
# raw=true # output to raw text without color
LabelColor=%s
TextColor=%s
EgColor=%s
`

func init() {
	var (
		fd    *os.File
		err   error
		first bool
	)
	_, e := os.Stat(inipath)
	if os.IsNotExist(e) {
		fd, err = os.Create(inipath)
		var DefaultIni = fmt.Sprintf(tpl, DefaultGoogleAPI, DefaultCibaAPI, DefaultYoudaoAPI, DefaultLabelColor, DefaultTextColor, DefaultEgColor)

		_, err := io.WriteString(fd, DefaultIni)
		err = fd.Close()
		if err != nil {
			log.Fatal(err)
		}
		first = true
	}
	c, err := ini.Load(inipath)
	Cfg = Config{c}
	if err != nil {
		log.Fatal(err)
	}
	Cfg.BlockMode = first // Block when first
}
