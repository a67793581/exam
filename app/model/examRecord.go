package model

type ExamRecord struct {
	Base
	Code        string
	Key         string
	ExamTime    int32
	Achievement int32
	CourseID    int32
	StudentID   int32
}
