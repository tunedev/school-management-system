package service

import (
	"fmt"
	"strconv"

	"github.com/tunedev/school-management-system/internal/models"
	"github.com/tunedev/school-management-system/internal/repository"
)

type CourseService interface {
    CreateCourse(course *models.Course)  error
    UpdateCourse(courseId int, course *models.Course) error
    GetCourseById(courseId string) (*models.Course, error)
    GetAllCourses() ([]*models.Course, error)
}

type courseService struct {
    repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
    return &courseService{repo: repo}
}

func (s *courseService) CreateCourse(course *models.Course) error {
	err := s.repo.CreateCourse(course)
    if err != nil {
        fmt.Println("Error occured while creating course ======>>>>>", err)
        return fmt.Errorf("internal server Error")
    }
    return nil
}

func (s *courseService) UpdateCourse(courseId int, course *models.Course) error {
	 err := s.repo.UpdateCourse(courseId, course)
    if err != nil {
        fmt.Println("Error occured while creating course ======>>>>>", err)
        return fmt.Errorf("internal server Error")
    }
    return nil
}

func (s *courseService) GetCourseById(id string) (*models.Course, error) {
   courseId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid course ID")
	}
    
	course, err := s.repo.GetCourseById(courseId)
    if err != nil {
        fmt.Println("Error occured while creating course ======>>>>>", err)
        return course, fmt.Errorf("internal server Error")
    }
    return course, nil
}

func (s *courseService) GetAllCourses() ([]*models.Course, error) {
	courses, err := s.repo.GetAllCourses()
    if err != nil {
        fmt.Println("Error occured while retrieving all courses ======>>>>>", err)
        return courses, fmt.Errorf("internal server Error")
    }
    return courses, nil
}