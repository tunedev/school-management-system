package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tunedev/school-management-system/internal/handler"
	"github.com/tunedev/school-management-system/internal/repository"
	"github.com/tunedev/school-management-system/internal/service"
	"github.com/tunedev/school-management-system/pkg/database"
)

func setupRouter(db *database.Database) *gin.Engine {
	r := gin.Default()

	// Initialize handlers
	// studentRepo := repository.NewStudentRepository(db)
	// studentService := service.NewStudentService(studentRepo)
	// studentHandler := handler.NewStudentHandler(studentService)

	courseRepo := repository.NewCourseRepository(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	studentRepo := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)

	// Define routes
	v1 := r.Group("/api/v1")
	{
		// Create a course
		v1.POST("/courses", courseHandler.CreateCourse)
		v1.PATCH("/courses/:id", courseHandler.UpdateCourse)
		v1.GET("/courses/:id", courseHandler.GetCourseById)
		v1.GET("/courses", courseHandler.GetAllCourses)
		v1.POST("/students", studentHandler.CreateStudent)
			v1.PATCH("/students/:id", studentHandler.UpdateStudent)
		v1.GET("/students/:id", studentHandler.GetStudentById)
		v1.GET("/students", studentHandler.GetAllStudents)
	}

	return r
}
