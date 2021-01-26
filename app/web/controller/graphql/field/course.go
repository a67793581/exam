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
		Description: "课程",
		Name:        "Course",
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
	},
)

func CourseList() *graphql.Field {
	return &graphql.Field{
		Description: "课程列表",
		Type:        graphql.NewList(&CourseObject),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result []model.Course
			id, ok := p.Args["id"].(int)
			db := mysql.GetIns().Model(&result)
			if ok {
				db.Where("id = ?", id)
			}
			db.Find(&result)
			return result, nil
		},
	}
}

func CourseShow() *graphql.Field {
	return &graphql.Field{
		Description: "课程",
		Type:        &CourseObject,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.Course
			id, ok := p.Args["id"].(int)
			db := mysql.GetIns().Model(&result)
			if ok {
				db.Where("id = ?", id)
			}
			db.First(&result)
			return result, nil
		},
	}
}
