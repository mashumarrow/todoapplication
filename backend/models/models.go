package models

type User struct {
	UserID string `gorm:"column:UserID;primaryKey;autoIncrement"`
	Name   string `gorm:"column:Name"`
}

type NewTodo struct {
	UserID    uint   `gorm:"column:UserID;primaryKey;autoIncrement"`
	Title     string `gorm:"column:Title"`
	SubjectID uint   `gorm:"column:SubjectName"`
}
type Subject struct {
	SubjectID   uint   `gorm:"column:SubjectID;primaryKey;autoIncrement"`
	SubjectName string `gorm:"column:SubjectName"`
}

type Classroom struct {
	ClassroomID   uint   `gorm:"column:ClassroomID;primaryKey;autoIncrement"`
	ClassroomName string `gorm:"column:ClassroomName"`
}

type Schedule struct {
	UserID      uint      `gorm:"column:UserID;primaryKey"`
	DayOfWeek   string    `gorm:"column:DayOfWeek"`
	Period      int       `gorm:"column:Period"`
	SubjectID   uint     `gorm:"column:SubjectID"`
	ClassroomID  uint    `gorm:"column:ClassroomID"`
}

type Todo struct {
	SubjectID uint    `gorm:"column:SubjectID"`
	Title     string  `gorm:"column:Title"`
	Completed bool    `gorm:"column:Completed"`
	UserID    uint    `gorm:"column:UserID"`
	
}
