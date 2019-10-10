package actions

import (
	"auth/models"
	"auth/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	if hasSession := pkg.HasSession(c); hasSession == true {
		c.String(http.StatusOK, "用户已登录")
		return
	}

	user := models.UserDetailByEmail(email)

	if err := pkg.Compare(user.Password, password); err != nil {
		c.String(401, "密码错误")
		return
	}

	pkg.SaveAuthSession(c, user.ID)

	c.Redirect(http.StatusMovedPermanently, "/page/welcome")
}

func Logout(c *gin.Context) {
	if hasSession := pkg.HasSession(c); hasSession != false {
		c.String(http.StatusUnauthorized, "用户未登录")
		return
	}
	pkg.ClearAuthSession(c)

	c.Redirect(http.StatusMovedPermanently, "/page/welcome")
}

func Register(c *gin.Context) {
	var user models.User
	user.Name = c.Request.FormValue("name")
	user.Email = c.Request.FormValue("email")

	if hasSession := pkg.HasSession(c); hasSession == true {
		c.String(200, "用户已登陆")
		return
	}

	if existUser := models.UserDetailByEmail(user.Email); existUser.ID != 0 {
		c.String(200, "邮箱已存在")
		return
	}

	if c.Request.FormValue("password") != c.Request.FormValue("password_confirmation") {
		c.String(200, "密码不一致")
		return
	}

	if pwd, err := pkg.Encrypt(c.Request.FormValue("password")); err == nil {
		user.Password = pwd
	}

	models.AddUser(&user)

	pkg.SaveAuthSession(c, user.ID)

	c.Redirect(http.StatusMovedPermanently, "/page/welcome")

}

func Me(c *gin.Context) {
	currentUser := c.MustGet("userId").(uint)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": currentUser,
	})
}

// Index index
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "page/login.html", gin.H{
		"data":    "Main website",
		"session": pkg.GetUserSession(c),
	})
}

// Index index
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "page/register.html", gin.H{
		"data":    "Main website",
		"session": pkg.GetUserSession(c),
	})
}