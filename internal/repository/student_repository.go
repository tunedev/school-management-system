package repository

import (
	"fmt"

	"github.com/tunedev/school-management-system/internal/models"
	"github.com/tunedev/school-management-system/pkg/database"
)

type StudentRepository interface {
	CreateStudent(student *models.Student) (int64, error)
	UpdateStudent(studentId int, student *models.Student) error
	GetStudentById(studentId int) (*models.Student, error)
	GetAllStudents() ([]*models.Student, error)
}

type studentRepository struct {
	DB *database.Database
}

func NewStudentRepository(db *database.Database) StudentRepository {
	return &studentRepository{DB: db}
}

func (r *studentRepository) CreateStudent(student *models.Student) (int64, error) {
	id, err := r.DB.CreateStudent(student)
	if err != nil {
		return 0, fmt.Errorf("failed to create student: %v", err)
	}
	return id, nil
}

func (r *studentRepository) UpdateStudent(id int, student *models.Student) error {
	 err := r.DB.UpdateStudent(id, student)
	if err != nil {
		return fmt.Errorf("failed to update student: %v", err)
	}
	return nil
}

func (r *studentRepository) GetStudentById(id int) (*models.Student, error) {
	student, err := r.DB.RetrieveStudentById(id)
	if err != nil {
		return student, fmt.Errorf("failed to get student by id: %v", err)
	}
	return student, nil
}

func (r *studentRepository) GetAllStudents() ([]*models.Student, error) {
	students, err := r.DB.RetrieveStudents()
	if err != nil {
		return students, fmt.Errorf("failed to get students: %v", err)
	}
	return students, nil
}
