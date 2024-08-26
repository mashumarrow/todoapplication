package models

import "gorm.io/gorm"
type User struct {
    gorm.Model
	UserID string `json:"userId"`  // UserIDをuint型に変更
    Name  string
    Email string
    Todos []Todo
}

type NewTodo struct {
	UserID    uint   `json:"userId"`
    Title     string `json:"title"`
    SubjectID uint   `json:"subjectId"`
}
type Subject struct {
    SubjectID   uint   `gorm:"primaryKey"`
    SubjectName string 
}

type Classroom struct {
    ClassroomID   uint   `gorm:"primaryKey"`
    ClassroomName string 
}

type Schedule struct {
    UserID         uint   `gorm:"primaryKey"`
    SubjectID  uint   `gorm:"not null"`
    ClassroomID uint  `gorm:"not null"`
    DayOfWeek  string `gorm:"type:enum('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday');not null"`
    Period     int    `gorm:"not null"`
    Subject    Subject
    Classroom  Classroom
}

type Todo struct {
	gorm.Model
    SubjectID uint   `gorm:"not null"`
    Title     string 
    Completed bool   `gorm:"default:false"`
    Subject   Subject
    UserID    uint   `json:"userId"`
    User      User
}
