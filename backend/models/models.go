package models


type User struct {
	UserID string   `json:"userid"`  
    Name  string    `json:"name"`
    Email string    `json:"email"`
    Todos []Todo    `json:"todos"`
}

type NewTodo struct {
	UserID    uint   `json:"userid"`
    Title     string `json:"title"`
    SubjectID uint   `json:"subjectId"`
}
type Subject struct {
    SubjectID   uint  `json:"subjectid"`  
    SubjectName string `json:"subjectName"`
}

type Classroom struct {
    ClassroomID   uint   `gorm:"column:ClassroomID;primaryKey;autoIncrement"`
        ClassroomName string   `gorm:"column:ClassroomName"`
}

type Schedule struct {
    UserID         uint   `json:"userid"` 
    SubjectID  uint    `json:"subjectid"`
    ClassroomID uint  `json:"classroomid"`
    DayOfWeek  string  `json:"dayOfWeek"`
    Period     int    `json:"period"`
    Subject    Subject `json:"subject"`
    Classroom  Classroom `json:"classroom"`
}

type Todo struct {

    SubjectID uint  `json:"subjectId"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
    Subject   Subject `json:"subject"`
    UserID    uint   `json:"userId"`
    User      User   `json:"user"`
}
