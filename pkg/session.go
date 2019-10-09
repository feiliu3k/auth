package pkg

import (
	"auth/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

const KEY = "AEN233"

// 使用 Cookie 保存 session
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("SAMPLE", store)
}

// session middleware
func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userId")
		if sessionValue == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Set("userId", sessionValue.(uint))
		c.Next()
		return
	}
}

// register and login should save the sessions
func SaveAuthSession(c *gin.Context, id uint)  {
	session := sessions.Default(c)
	session.Set("userId", id)
	session.Save()
}

// when logout clear the session
func ClearAuthSession(c *gin.Context)  {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("userId"); sessionValue ==nil {
		return false
	}
	return true
}

func GetSessionUserId(c *gin.Context) uint {
	session := sessions.Default(c)
	sessionValue := session.Get("userId")
	if sessionValue == nil {
		return 0
	}
	return  sessionValue.(uint)
}

func GetUserSession(c *gin.Context) map[string]interface{} {
	HasSession := HasSession(c)
	userName := ""
	if HasSession {
		userId := GetSessionUserId(c)
		userName = models.UserDetail(userId).Name
	}
	data := make(map[string]interface{})
	data["hasSession"] = HasSession
	data["userName"] = userName
	return data
}