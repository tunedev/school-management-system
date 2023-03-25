package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/tunedev/school-management-system/internal/models"
)

func (d *Database) CreateCourse(course *models.Course) (int64, error) {
	stmt, err := d.DB.Prepare("INSERT INTO courses (name, description) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	resp, err := stmt.Exec(course.Name, course.Description)
	if err != nil {
		return 0, fmt.Errorf("failed to create course: %v", err)
	}

	id, _ := resp.LastInsertId()

	return id, nil
}

func (d *Database) UpdateCourse(courseId int, course *models.Course) (int64, error) {
    resp, err := d.DB.Exec("UPDATE courses SET name=?, description=?, updated_at=? WHERE id=?", course.Name, course.Description, time.Now(), courseId)
    if err != nil {
        return 0, fmt.Errorf("failed to update course: %v", err)
    }
   id, _ := resp.LastInsertId()

	return id, nil
}

// RetrieveCourse retrieves a course from the database given a course ID
func (d *Database) RetrieveCourseById(courseID int) (*models.Course, error) {
    var course models.Course
    row := d.DB.QueryRow("SELECT id, name, description, created_at, updated_at FROM courses WHERE id = ?", courseID)
    err := row.Scan(&course.ID, &course.Name, &course.Description, &course.CreatedAt, &course.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("course not found")
        }
        return nil, fmt.Errorf("failed to retrieve course: %v", err)
    }
    return &course, nil
}

// RetrieveCourses retrieves all courses from the database
func (d *Database) RetrieveCourses() ([]*models.Course, error) {
    var courses []*models.Course
    rows, err := d.DB.Query("SELECT id, name, description, created_at, updated_at FROM courses")
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve courses: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var course models.Course
        err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CreatedAt, &course.UpdatedAt)
        if err != nil {
            return nil, fmt.Errorf("failed to retrieve courses: %v", err)
        }
        courses = append(courses, &course)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("failed to retrieve courses: %v", err)
    }

    return courses, nil
}