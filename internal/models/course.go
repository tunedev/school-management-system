package models

import "time"

// TODO: Define a struct for a course
type Course struct {
    ID          uint64 `json:"id,omitempty"`
    Name        string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    CreatedAt   time.Time `json:"created_at,omitempty"`
    UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

