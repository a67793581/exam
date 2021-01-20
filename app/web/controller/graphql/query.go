package graphql

import (
	"exam/app/web/controller/graphql/field"
	"github.com/graphql-go/graphql"
)

func newQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"Users":       field.NewUsers(),
			"User":        field.NewUser(),
			"ExamRecords": field.NewExamRecords(),
			"ExamRecord":  field.NewExamRecord(),
		},
	})
}
