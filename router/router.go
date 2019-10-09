package router

import (
	"auth/apis"
	"auth/pkg"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/index", apis.UserIndex)
	}

	// use session router
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