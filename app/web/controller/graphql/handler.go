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
						"examRecords": field.ExamRecordList(),
						"examRecord":  field.ExamRecordShow(),
						"students":    field.StudentList(),
						"student":     field.StudentShow(),
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
						"examRecords": field.ExamRecordList(),
						"examRecord":  field.ExamRecordShow(),
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
