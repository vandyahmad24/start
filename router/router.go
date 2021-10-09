package router

import (
	"test-start/config"
	"test-start/handler"
	"test-start/models/students"

	"github.com/gin-gonic/gin"
)

func Router() {
	studentsRepository := students.NewRepository(config.DB)
	studentsService := students.NewService(studentsRepository)
	studentsHandler := handler.NewUserHandler(studentsService)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Created By Vandy Ahmad",
		})
	})
	api := r.Group("student")
	api.POST("", studentsHandler.StoreStudent)
	api.PUT("/:id", studentsHandler.PutStudent)
	api.GET("/:id", studentsHandler.GetSudent)
	api.DELETE("/:id", studentsHandler.DeleteStudent)
	r.Run(config.GetEnvVariable("SERVER_HOST"))
}
