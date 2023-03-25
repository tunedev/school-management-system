package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/tunedev/school-management-system/internal/models"
)

func (d *Database) CreateStudent(student *models.Student) (int64, error) {
	stmt, err := d.DB.Prepare("INSERT INTO students (name, email) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	resp, err := stmt.Exec(student.Name, student.Email)
	if err != nil {
		return 0, fmt.Errorf("failed to create student: %v", err)
	}

	id, _ := resp.LastInsertId()

	return id, nil
}

func (d *Database) UpdateStudent(studentId int, student *models.Student) error {
    _, err := d.DB.Exec("UPDATE student SET name=?, email=?, updated_at=? WHERE id=?", student.Name, student.Email, time.Now(), studentId)
    if err != nil {
        return fmt.Errorf("failed to update student: %v", err)
    }
    return nil
}

// RetrieveCourse retrieves a student from the database given a student ID
func (d *Database) RetrieveStudentById(studentID int) (*models.Student, error) {
    var student models.Student
    row := d.DB.QueryRow("SELECT id, name, email, created_at, updated_at FROM students WHERE id = ?", studentID)
    err := row.Scan(&student.ID, &student.Name, &student.Email, &student.CreatedAt, &student.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("student not found")
        }
        return nil, fmt.Errorf("failed to retrieve student: %v", err)
    }
    return &student, nil
}

// RetrieveCourses retrieves all students from the database
func (d *Database) RetrieveStudents() ([]*models.Student, error) {
    var students []*models.Student
    rows, err := d.DB.Query("SELECT id, name, email, created_at, updated_at FROM students")
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve students: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var student models.Student
        err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.CreatedAt, &student.UpdatedAt)
        if err != nil {
            return nil, fmt.Errorf("failed to retrieve students: %v", err)
        }
        students = append(students, &student)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("failed to retrieve students: %v", err)
    }

    return students, nil
}



