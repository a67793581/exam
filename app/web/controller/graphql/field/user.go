package field

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"exam/app/web/controller/graphql/object"
	"github.com/graphql-go/graphql"
)

func NewUsers() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(object.User),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result []*model.User
			mysql.GetIns().Find(&result)
			return result, nil
		},
		Description: "用户列表",
	}
}
func NewUser() *graphql.Field {
	return &graphql.Field{
		Type: object.User,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.User
			mysql.GetIns().First(&result)
			return result, nil
		},
		Description: "用户",
	}
}
