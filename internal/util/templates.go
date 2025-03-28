package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Carrega os templates corretamente no engine do Gin
func CarregarTemplates(r *gin.Engine) {
	r.LoadHTMLGlob("web/templates/*")
}

// Executa o template usando o sistema do Gin
func ExecutarTemplate(c *gin.Context, tmpl string, data interface{}) {
	c.HTML(http.StatusOK, tmpl, data)
}
