package field

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"github.com/graphql-go/graphql"
)

type Course struct {
	Object *graphql.Object
}

var CourseObject = *graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Course",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					Course, ok := (p.Source).(model.Course)
					if !ok {
						panic("传入参数不是model.Course变量")
					}
					return Course.ID, nil
				}},
			"created_at": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					Course, ok := (p.Source).(model.Course)
					if !ok {
						panic("传入参数不是model.Course变量")
					}
					return Course.CreatedAt, nil
				},
			},
			"updated_at": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					Course, ok := (p.Source).(model.Course)
					if !ok {
						panic("传入参数不是model.Course变量")
					}
					return Course.UpdatedAt, nil
				},
			},
			"deleted_at": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					Course, ok := (p.Source).(model.Course)
					if !ok {
						panic("传入参数不是model.Course变量")
					}
					return Course.DeletedAt, nil
				},
			},
			"name": &graphql.Field{Type: graphql.String},
		},
		Description: "课程",
	},
)

func CourseList() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(&CourseObject),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result []model.Course
			mysql.GetIns().Find(&result)
			return result, nil
		},
		Description: "课程列表",
	}
}

func CourseShow() *graphql.Field {
	return &graphql.Field{
		Type: &CourseObject,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.Course
			mysql.GetIns().First(&result)
			return result, nil
		},
		Description: "课程",
	}
}
