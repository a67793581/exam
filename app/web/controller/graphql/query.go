package graphql

import (
	"exam/app/web/controller/graphql/field"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

func newQuery(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"ExamRecords": field.ExamRecords(db),
			"ExamRecord":  field.ExamRecord(db),
		},
	})
}
