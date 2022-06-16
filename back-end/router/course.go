package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/controller"
	"github.com/gin-gonic/gin"
)

func CourseRouter(Router *gin.RouterGroup) {
	CourseRouter := Router.Group("courses")
	{
		CourseRouter.GET("/:course_name/assessments/:assessment_name/exam", controller.Exam_Handler)
		// CourseRouter.POST("/token", controller.Authtoken_Handler)
	}
}
