package repository

import (
	"fmt"

	"github.com/tunedev/school-management-system/internal/models"
	"github.com/tunedev/school-management-system/pkg/database"
)

type CourseRepository interface {
	CreateCourse(course *models.Course) error
	UpdateCourse(courseId int, course *models.Course)  error
	GetCourseById(courseId int) (*models.Course, error)
	GetAllCourses() ([]*models.Course, error)
}

type courseRepository struct {
	DB *database.Database
}

func NewCourseRepository(db *database.Database) CourseRepository {
	return &courseRepository{DB: db}
}

func (r *courseRepository) CreateCourse(course *models.Course)  error {
	id, err := r.DB.CreateCourse(course)
	if err != nil {
		return fmt.Errorf("failed to create course: %v", err)
	}
	newCourse, _ := r.GetCourseById(int(id))
	*course = *newCourse
	return nil
}

func (r *courseRepository) UpdateCourse(coursId int, course *models.Course)  error {
	id, err := r.DB.UpdateCourse(coursId, course)
	if err != nil {
		return fmt.Errorf("failed to update course: %v", err)
	}
	newCourse, _ := r.GetCourseById(int(id))
	*course = *newCourse
	return nil
}

func (r *courseRepository) GetCourseById(id int) (*models.Course, error) {
	course, err := r.DB.RetrieveCourseById(id)
	if err != nil {
		return course, fmt.Errorf("failed to retrieve course: %v", err)
	}
	return course, nil
}

func (r *courseRepository) GetAllCourses() ([]*models.Course, error) {
	courses, err := r.DB.RetrieveCourses()
	if err != nil {
		return courses, fmt.Errorf("failed to retrieve course: %v", err)
	}
	return courses, nil
}
