package routes

import (
	"apallis/portfolio-front/handler"

	"github.com/gin-gonic/gin"
)


var homeHandler = handler.HomeHandler{}

func Init(router *gin.Engine) {
    router.GET("/", homeHandler.Show)
}
