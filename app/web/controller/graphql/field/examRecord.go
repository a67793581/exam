package field

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"exam/app/web/controller/graphql/object"
	"github.com/graphql-go/graphql"
)

func NewExamRecords() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(object.ExamRecord),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result []*model.ExamRecord
			mysql.GetIns().Find(&result)
			return result, nil
		},
		Description: "考试记录列表",
	}
}

func NewExamRecord() *graphql.Field {
	return &graphql.Field{
		Type: object.ExamRecord,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.ExamRecord
			mysql.GetIns().First(&result)
			return result, nil
		},
		Description: "考试记录",
	}
}
