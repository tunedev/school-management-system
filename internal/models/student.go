package models

import "time"

// TODO: Define a struct for a course
type Student struct {
    ID          uint64 `json:"id,omitempty"`
    Name        string `json:"name,omitempty"`
    Email string `json:"email,omitempty"`
    CreatedAt   time.Time `json:"created_at,omitempty"`
    UpdatedAt   time.Time `json:"updated_at,omitempty"`
}