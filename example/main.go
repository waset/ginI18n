package main

import (
	"log"
	"net/http"
	"path"

	i18n "github.com/waset/ginI18n"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// apply i18n middleware
	router.Use(i18n.Localizer(&i18n.Options{
		DefaultLang:  "zh-CN",                          // default language
		SupportLangs: "zh-CN,en-US",                    // list of supported languages ​​(must include default language)
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
