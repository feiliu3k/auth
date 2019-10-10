package routes

import (
	"auth/actions"
	"auth/apis"
	"auth/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.Default()

	pageGroup(r)

	user := r.Group("/user")
	{
		user.GET("/index", apis.UserIndex)
	}

	// use session routes
	sr := r.Group("/", pkg.EnableCookieSession())
	{
		sr.GET("/welcome", apis.Welcome)
		sr.GET("/login", apis.Login)
		sr.GET("/register", apis.Register)
		sr.GET("/logout", apis.Logout)

		authorized  := sr.Group("/auth", pkg.AuthSessionMiddle())
		{
			authorized.GET("/me", apis.Me)
		}
	}

	return r
}

func pageGroup(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("views/**/*")
	sr :=r.Group("/page", pkg.EnableCookieSession())
	{
		sr.StaticFS("/assets", http.Dir("assets"))
		sr.GET("/welcome", actions.Welcome)
		sr.GET("/login_page", actions.LoginPage)
		sr.GET("/register_page", actions.RegisterPage)
		sr.POST("/login", actions.Login)
		sr.POST("/register", actions.Register)
		sr.POST("/logout", actions.Logout)

		authorized := sr.Group("/auth", pkg.AuthSessionMiddle())
		{
			authorized.GET("/me", actions.Me)
		}
	}

	return r
}