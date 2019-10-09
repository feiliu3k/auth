package apis

import (
	"auth/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserIndex(c *gin.Context)  {


	//models.DB.Create(&models.User{Name: "admin", Password: "123456", Email: "admin@admin.com", Gender: "male"})

	var user models.User
	result := models.DB.Take(&user).Value
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}
