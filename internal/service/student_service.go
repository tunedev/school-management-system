package service

import (
	"fmt"
	"strconv"

	"github.com/tunedev/school-management-system/internal/models"
	"github.com/tunedev/school-management-system/internal/repository"
)

type StudentService interface {
    CreateStudent(student *models.Student) (int64, error)
    UpdateStudent(studentId int, student *models.Student) error
    GetStudentById(studentId string) (*models.Student, error)
    GetStudents() ([]*models.Student, error)
}

type studentService struct {
    repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
    return &studentService{repo: repo}
}

func (s *studentService) CreateStudent(student *models.Student) (int64, error) {
	 id, err := s.repo.CreateStudent(student)
    if err != nil {
        fmt.Println("Error occured while creating student ======>>>>>", err)
        return 0, fmt.Errorf("internal server Error")
    }
    return id, nil
}

func (s *studentService) UpdateStudent(id int, student *models.Student) error {
	err := s.repo.UpdateStudent(id, student)
    if err != nil {
        fmt.Println("error occured while updating student ======>>>>>", err)
        return fmt.Errorf("internal server Error")
    }
    return nil
}

func (s *studentService) GetStudentById(id string) (*models.Student, error) {
    studentId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid student ID")
	}
	student, err := s.repo.GetStudentById(studentId)
    if err != nil {
        fmt.Println("error occured while retriving student By Id ======>>>>>", err)
        return student, fmt.Errorf("internal server Error")
    }
    return student, nil
}

func (s *studentService) GetStudents() ([]*models.Student, error) {
	students, err := s.repo.GetAllStudents()
    if err != nil {
        fmt.Println("error occured while retriving all students ======>>>>>", err)
        return students, fmt.Errorf("internal server Error")
    }
    return students, nil
}