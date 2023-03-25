package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tunedev/school-management-system/internal/models"
	"github.com/tunedev/school-management-system/internal/service"
)

type CourseHandler struct {
    service service.CourseService
}

func NewCourseHandler(service service.CourseService) *CourseHandler {
    return &CourseHandler{service: service}
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {
    var course models.Course
    if err := c.ShouldBindJSON(&course); err != nil {
				fmt.Println("Error ====>>>", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
        return
    }

		err := h.service.CreateCourse(&course)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create course"})
        return
    }

    c.JSON( http.StatusCreated, course)
}

func (h *CourseHandler) UpdateCourse(c *gin.Context) {
  id := c.Param("id")
	var course models.Course

	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}
	err = h.service.UpdateCourse(ID, &course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update course"})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) GetCourseById(c *gin.Context) {
  id := c.Param("id")

  course, err := h.service.GetCourseById(id) 

  if err != nil {
    if err.Error() == "invalid course ID"{
      c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
      return
    }
    c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrive course"})
		return
  }

  if course == nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		return
  }

	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) GetAllCourses(c *gin.Context) {
  courses, err := h.service.GetAllCourses() 

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrive course"})
		return
  }

	c.JSON(http.StatusOK, courses)
}