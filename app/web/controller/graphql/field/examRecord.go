package field

import (
	"exam/app/model"
	"exam/app/web/controller/graphql/object"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

func ExamRecords(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(object.ExamRecord),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result []*model.ExamRecord
			db.Find(&result)
			return result, nil
		},
		Description: "考试记录列表",
	}
}

func ExamRecord(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: object.ExamRecord,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result *model.ExamRecord
			db.First(&result)
			return result, nil
		},
		Description: "考试记录",
	}
}
