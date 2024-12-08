package models

type User struct {
	UserID   string `gorm:"column:UserID;primaryKey;autoIncrement"`
	Name     string `gorm:"column:Name"`
	Password string `gorm:"column:Password"`
	Token    string `gorm:"-"`
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
	ScheduleID uint   `gorm:"primaryKey;autoIncrement;column:ScheduleID"`
	UserID     uint   `gorm:"column:UserID"`
	DayOfWeek  string `gorm:"column:DayOfWeek"`
	Period     int    `gorm:"column:Period"`

	SubjectName string `gorm:"column:SubjectName"`

	ClassroomName string `gorm:"column:ClassroomName"`
}

type Todo struct {
	TodoID    uint   `gorm:"primaryKey;autoIncrement;column:TodoID"`
	Title     string `gorm:"column:Title"`
	Completed bool   `gorm:"column:Completed"`
	UserID    uint   `gorm:"column:UserID"`
	//DayOfWeek string `gorm:"column:DayOfWeek"`
	Period        int    `gorm:"column:Period"`
	Subjectname   string `gorm:"column:Subjectname"`
	Classroomname string `gorm:"column:Classroomname"`
}

type NewTodo struct {
	Title  string `gorm:"column:Title"`
	UserID uint   `gorm:"column:UserID"`
	//DayOfWeek string `gorm:"column:DayOfWeek"`
	Period        int    `gorm:"column:Period"`
	Subjectname   string `gorm:"column:Subjectname"`
	Classroomname string `gorm:"column:Classroomname"`
}
