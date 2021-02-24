package model

type ExamRecord struct {
	Base
	Code        string `gorm:"index;comment:考试批次"`
	Key         string `gorm:"index;comment:考试唯一编号"`
	ExamTime    int32  `gorm:"index;comment:考试时间"`
	Achievement int32  `gorm:"index;comment:成绩"`
	CourseID    int32  `gorm:"index;comment:课程id"`
	StudentID   int32  `gorm:"index;comment:学生id"`
}
