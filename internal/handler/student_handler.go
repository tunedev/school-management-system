package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tunedev/school-management-system/internal/models"
	"github.com/tunedev/school-management-system/internal/service"
)

type StudentHandler struct {
    service service.StudentService
}

func NewStudentHandler(service service.StudentService) *StudentHandler {
    return &StudentHandler{service: service}
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
    var student models.Student
    if err := c.ShouldBindJSON(&student); err != nil {
				fmt.Println("Error ====>>>", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
        return
    }

		id, err := h.service.CreateStudent(&student)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create course"})
        return
    }
		student.ID = uint64(id)

    c.JSON( http.StatusCreated, student)
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
  id := c.Param("id")
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student ID"})
		return
	}

	if err := h.service.UpdateStudent(ID, &student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update student"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) GetStudentById(c *gin.Context) {
  id := c.Param("id")

  student, err := h.service.GetStudentById(id) 

  if err != nil {
    if err.Error() == "invalid student ID"{
      c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student ID"})
      return
    }
    c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrive student"})
		return
  }

  if student == nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
  }

	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) GetAllStudents(c *gin.Context) {
  students, err := h.service.GetStudents() 

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrive students"})
		return
  }

	c.JSON(http.StatusOK, students)
}