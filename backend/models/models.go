package models

type User struct {
	UserID string `gorm:"column:UserID;primaryKey;autoIncrement"`
	Name   string `gorm:"column:Name"`
	Password   string `gorm:"column:Password"`
	Token    string `gorm:"-"`
}

type NewTodo struct {
	UserID    uint   `gorm:"column:UserID;primaryKey;autoIncrement"`
	Title     string `gorm:"column:Title"`
	SubjectID uint   `gorm:"column:SubjectName"`
}
type Subject struct {
	SubjectID   uint   `gorm:"column:SubjectID;primaryKey;autoIncrement"`
	SubjectName string `gorm:"column:SubjectName"`
	Todos       []Todo `gorm:"foreignKey:SubjectID"`
}

type Classroom struct {
	ClassroomID   uint   `gorm:"column:ClassroomID;primaryKey;autoIncrement"`
	ClassroomName string `gorm:"column:ClassroomName"`
}

type Schedule struct {
	UserID      uint   `gorm:"column:UserID"`
	DayOfWeek   string `gorm:"column:DayOfWeek"`
	Period      int    `gorm:"column:Period"`
	SubjectID   uint   `gorm:"column:SubjectID"`
	SubjectName string `gorm:"column:SubjectName"`
	ClassroomID uint   `gorm:"column:ClassroomID"`
	ClassroomName string `gorm:"column:ClassroomName"`
}

type Todo struct {
	SubjectID uint    `gorm:"column:SubjectID"`
	Title     string  `gorm:"column:Title"`
	Completed bool    `gorm:"column:Completed"`
	UserID    uint    `gorm:"column:UserID"`
	User      User    `gorm:"foreignKey:UserID"`
	Subject   Subject `gorm:"foreignKey:SubjectID"`
}
