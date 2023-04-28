# 简单的 gin 多语言支持中间件

[English](README.md) | 简体中文

## 鸣谢
[fishjar/gin-i18n](https://github.com/fishjar/gin-i18n)

## 使用

```go
package main

import (
	"log"
	"net/http"
	"path"

	i18n "github.com/waset/ginI18n"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// apply i18n middleware
	router.Use(i18n.Localizer(&i18n.Options{
		DefaultLang:  "zh-Hans",        // default language
		SupportLangs: "zh-Hans,en-US",                    // list of supported languages ​​(must include default language)
		FilePath:     path.Join("example", "localize"), // multilingual file directory
	}))

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, i18n.Msg(c, "welcome"))
	})

	router.GET("/:name", func(c *gin.Context) {
		c.String(http.StatusOK, i18n.Msg(c, "hello_world", c.Param("name")))
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
```

```yaml
# example/localize/zh-CN.yml
welcome: 欢迎！
hello_world: 你好 %s!
```

```yaml
# example/localize/en-US.yml
welcome: welcome!
hello_world: hello %s!
```

```sh
# 检查效果
go run example/main.go

curl http://127.0.0.1:8080/ -H 'accept-language: zh-CN'             # 欢迎！
curl http://127.0.0.1:8080/ -H 'accept-language: en-US'             # welcome!

curl http://127.0.0.1:8080/ -H 'accept-language: zh-CN,en-US;q=0.9' # 欢迎！
curl http://127.0.0.1:8080/ -H 'accept-language: zh-CN;q=0.9,en-US' # welcome!

curl http://127.0.0.1:8080/ -H 'accept-language: zh'                # 欢迎！
curl http://127.0.0.1:8080/ -H 'accept-language: en'                # welcome!

curl http://127.0.0.1:8080/gabe -H 'accept-language: zh-CN'         # 你好 gabe!
curl http://127.0.0.1:8080/gabe -H 'accept-language: en-US'         # hello gabe!
```

## 配置

<details>
<summary>i18n-ally</summary>
<br>

[i18n-ally](https://github.com/lokalise/i18n-ally)

```json
// .vscode/settings.json
{
    "i18n-ally.localesPaths": [
        "locales"
    ],
    "i18n-ally.sourceLanguage": "zh-Hans",
    "i18n-ally.enabledFrameworks": [
        "custom",
    ],
    "i18n-ally.sortKeys": true,
}

```

```yml
# .vscode/i18n-ally-custom-framework.yml

# An array of strings which contain Language Ids defined by VS Code
# You can check avaliable language ids here: https://code.visualstudio.com/docs/languages/overview#_language-id
languageIds:
  - go

# An array of RegExes to find the key usage. **The key should be captured in the first match group**.
# You should unescape RegEx strings in order to fit in the YAML file
# To help with this, you can use https://www.freeformatter.com/json-escape.html
usageMatchRegex:
  # The following example shows how to detect `t("your.i18n.keys")`
  # the `{key}` will be placed by a proper keypath matching regex,
  # you can ignore it and use your own matching rules as well
  - "(?:i18n\\.Msg\\(\\w+, *['\"]([\\w.-]+)['\"](?: *, *[\\w{}]+)*\\))"


# An array of strings containing refactor templates.
# The "$1" will be replaced by the keypath specified.
# Optional: uncomment the following two lines to use

refactorTemplates:
 - i18n.Msg(c, "$1")


# If set to true, only enables this custom framework (will disable all built-in frameworks)
monopoly: true

```
</details>
