# translator

命令行翻译工具
```
Usage of tl:
  -e string
    	engine (default "google")
  -s string
    	source language (default "auto")
  -t string
    	target language
Supported engines: youdao, google, ciba

```
| engine名 | 中文       | 备注     |
| ------- | -------- | ------ |
| google  | Google翻译 | ajax接口 |
| ciba    | 金山词霸     | ajax   |
| youdao  | 有道翻译     | ajax接口 |

通过[环境变量](https://golang.org/pkg/net/http/#ProxyFromEnvironment)`http_proxy` or `https_proxy`来代理请求

感谢 [翻译接口总结](https://juejin.im/post/5beaac9cf265da614a3a09a9)

![](https://i.imgur.com/JE0qi6h.png)
