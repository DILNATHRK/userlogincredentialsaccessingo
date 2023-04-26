package routes

import (
	"commercial-propfloor-users/users"

	"github.com/gin-gonic/gin"
)

func Routes(Routes *gin.Engine) {
	Routes.POST("/userlogin", users.AddUserdetails())
}
