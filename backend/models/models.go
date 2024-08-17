package models

//"gorm.io/gorm"

type NewTodo struct {
    Title     string `json:"title"`
    SubjectID uint   `json:"subjectId"`
}
type Subject struct {
    ID   uint   `gorm:"primaryKey"`
    Name string 
}

type Classroom struct {
    ID   uint   `gorm:"primaryKey"`
    Name string 
}

type Schedule struct {
    ID         uint   `gorm:"primaryKey"`
    SubjectID  uint   `gorm:"not null"`
    ClassroomID uint  `gorm:"not null"`
    DayOfWeek  string `gorm:"type:enum('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday');not null"`
    Period     int    `gorm:"not null"`
    Subject    Subject
    Classroom  Classroom
}

type Todo struct {
    ID        uint   `gorm:"primaryKey"`
    SubjectID uint   `gorm:"not null"`
    Title     string 
    Completed bool   `gorm:"default:false"`
    Subject   Subject
}
