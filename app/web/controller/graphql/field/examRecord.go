package field

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"github.com/graphql-go/graphql"
)

type ExamRecord struct {
	Object *graphql.Object
}

var ExamRecordObject = *graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ExamRecord",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.ID, nil
					}
					return 0, nil
				}},
			"created_at": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.CreatedAt, nil
					}
					return 0, nil
				},
			},
			"updated_at": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.UpdatedAt, nil
					}
					return 0, nil
				},
			},
			"deleted_at": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.DeletedAt, nil
					}
					return 0, nil
				},
			},
			"code": &graphql.Field{Type: graphql.String},
			"key":  &graphql.Field{Type: graphql.String},
			"exam_time": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.ExamTime, nil
					}
					return 0, nil
				},
			},
			"achievement": &graphql.Field{Type: graphql.Int},
			"course_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.CourseID, nil
					}
					return 0, nil
				},
			},
			"student_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.StudentID, nil
					}
					return 0, nil
				},
			},
			"student": StudentShow(),
			"course":  CourseShow(),
		},
		Description: "考试记录",
	},
)

func ExamRecordList() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(&ExamRecordObject),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result []model.ExamRecord
			mysql.GetIns().Find(&result)
			return result, nil
		},
		Description: "考试记录列表",
	}
}

func ExamRecordShow() *graphql.Field {
	return &graphql.Field{
		Type: &ExamRecordObject,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.ExamRecord
			mysql.GetIns().First(&result)
			return result, nil
		},
		Description: "考试记录",
	}
}
