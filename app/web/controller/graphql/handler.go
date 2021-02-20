package graphql

import (
	"exam/app/web/controller/graphql/field"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func Teacher() (*handler.Handler, error) {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Query",
					Fields: graphql.Fields{
						"examRecords":          field.ExamRecordList(),
						"examRecord":           field.ExamRecordShow(),
						"examRecordConnection": field.ExamRecordConnection(),
						"students":             field.StudentList(),
						"student":              field.StudentShow(),
						"studentConnection":    field.StudentConnection(),
						"courses":              field.CourseList(),
						"course":               field.CourseShow(),
						"courseConnection":     field.CourseConnection(),
					},
				},
			),
			Mutation: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Mutation",
					Fields: graphql.Fields{
						"examRecordDML": field.ExamRecordDML(),
						"studentDML":    field.StudentDML(),
						"courseDML":     field.CourseDML(),
					},
				},
			),
		},
	)
	if err != nil {
		return nil, err
	}
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}), nil
}

func Student() (*handler.Handler, error) {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Query",
					Fields: graphql.Fields{
						"examRecord": field.ExamRecordShow(),
					},
				},
			),
		},
	)
	if err != nil {
		return nil, err
	}
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}), nil
}
