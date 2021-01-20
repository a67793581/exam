package object

import "github.com/graphql-go/graphql"

var ExamRecord = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ExamRecord",
		Fields: graphql.Fields{
			"ID":          &graphql.Field{Type: graphql.ID},
			"created_at":  &graphql.Field{Type: graphql.Int},
			"updated_at":  &graphql.Field{Type: graphql.Int},
			"deleted_at":  &graphql.Field{Type: graphql.Int},
			"code":        &graphql.Field{Type: graphql.String},
			"key":         &graphql.Field{Type: graphql.String},
			"exam_time":   &graphql.Field{Type: graphql.Int},
			"achievement": &graphql.Field{Type: graphql.Int},
			"course_id":   &graphql.Field{Type: graphql.Int},
			"student_id":  &graphql.Field{Type: graphql.Int},
		},
		Description: "考试记录",
	},
)
