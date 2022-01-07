package app

import (
	"github.com/chpushpa/student/controllers"
)

func mapUrls() {
	router.POST("/student", controllers.StudentsController.Create)
	router.GET("/student/:id", controllers.StudentsController.Get)
}
