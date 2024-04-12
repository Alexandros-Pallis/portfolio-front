package main

import (
	"apallis/portfolio-front/renderer"
	"apallis/portfolio-front/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	ginHTMLRenderer := router.HTMLRender
	router.HTMLRender = &renderer.HTMLTemplRenderer{
		FallbackHtmlRenderer: ginHTMLRenderer,
	}
	router.SetTrustedProxies(nil)
	router.Static("/dist", "./dist")
	router.Static("/icons", "./assets/icon")
	router.StaticFile("/favicon.png", "./assets/favicon.png")
	router.StaticFile("/bio.pdf", "./assets/bio.pdf")
	routes.Init(router)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
    log.Fatal(autotls.Run(router, "alexandrospallis.dev", "www.alexandrospallis.dev"))
	// router.Run(":80")
}
