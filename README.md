# translator
[![Build Status](https://travis-ci.org/fengkx/translator.svg?branch=master)](https://travis-ci.org/fengkx/translator)       
多引擎命令行翻译工具   [PreBuild下载](https://github.com/fengkx/translator/releases)  
```
Usage of tl:
  -e string
    	engine (default "google")
  -raw
    	raw output without color escape
  -s string
    	source language (default "auto")
  -t string
    	target language
Supported engines: ciba, youdao, bdfanyi, google
Config ini path: /home/fengkx/.config

```
| engine名 | 中文       | 备注     |
| ------- | -------- | ------ |
| google  | Google翻译 | ajax接口 |
| ciba    | 金山词霸     | ajax接口   |
| youdao  | 有道翻译     | ajax接口 |
| bdfanyi  | 百度翻译     | 商业接口 |

通过[环境变量](https://golang.org/pkg/net/http/#ProxyFromEnvironment)`http_proxy` or `https_proxy`来代理请求

通过 ini 配置文件设置输出终端颜色，API host等参数。ini路径在Usage中有显示 文件名为`go-translator.ini`

默认ini内容
```ini
# Translator configuration
[google]
HOST=https://translate.googleapis.com/

[ciba]
HOST=http://fy.iciba.com/ajax.php

[youdao]
HOST=http://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule

[bdfanyi]
HOST=https://fanyi-api.baidu.com/api/trans/vip/translate
APPID=XXXXXXXXX
APPKEY=XXXXXXXXXX

[output]
# only support black red green yellow blue magenta cyan white
# raw=true # output to raw text without color
LabelColor=green
TextColor=white
EgColor=yellow

```

感谢 [翻译接口总结](https://juejin.im/post/5beaac9cf265da614a3a09a9)

![windows](https://i.imgur.com/urEOQbE.png)
![linux](https://i.imgur.com/PHQ6O4F.png)
