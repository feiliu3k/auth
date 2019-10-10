package actions

import (
	"auth/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "page/welcome.html", gin.H{
		"data": "Main website",
		"session": pkg.GetUserSession(c),
	})
}
