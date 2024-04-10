package handler

import (
	"apallis/portfolio-front/renderer"
	"apallis/portfolio-front/view/page"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct{}

func (h *HomeHandler) Show(c *gin.Context) {
	r := renderer.New(c, http.StatusOK, page.Home())
	c.Render(http.StatusOK, r)
}
