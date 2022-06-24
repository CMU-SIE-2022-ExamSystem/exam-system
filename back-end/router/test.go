package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/controller"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/gin-gonic/gin"
)

func TestRouter(Router *gin.RouterGroup) {
	TestRouter := Router.Group("test")
	{
		TestRouter.GET("/", func(c *gin.Context) {
			panic("123")
		})

		TestRouter.GET("/users", jwt.JWTAuth(), controller.GetUsers)
		TestRouter.GET("/login", controller.Login)
		TestRouter.GET("/refresh", jwt.JWTAuth(), jwt.UserRefreshHandler)
		TestRouter.GET("/cookie", controller.CookieTest)
		TestRouter.GET("/:course_name/check", jwt.JWTAuth(), jwt.Check_authlevel)
		TestRouter.GET("/f/:course/:assessment/:user_id/", controller.FolderTest)
	}
}
