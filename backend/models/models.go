package models

type User struct {
	UserID string `gorm:"column:UserID;primaryKey;autoIncrement"`
	Name   string `gorm:"column:Name"`
	Email  string `gorm:"column:Email"`
	Todos  []Todo `gorm:"column:Todos"`
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
	UserID      uint      `gorm:"column:UserID;primaryKey;autoIncrement"`
	DayOfWeek   string    `gorm:"column:DayOfWeek"`
	Period      int       `gorm:"column:Period"`
	Subject     Subject    `gorm:"foreignKey:SubjectID"`
	Classroom   Classroom `gorm:"foreignKey:ClassroomID"`
}

type Todo struct {
	SubjectID uint    `json:"subjectId"`
	Title     string  `json:"title"`
	Completed bool    `json:"completed"`
	Subject   Subject `json:"subject"`
	UserID    uint    `json:"userId"`
	User      User    `json:"user"`
}
