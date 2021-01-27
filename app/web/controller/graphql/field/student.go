package field

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"github.com/graphql-go/graphql"
)

type Student struct {
	Object *graphql.Object
}

var StudentObject = *graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Student",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					Student, ok := (p.Source).(model.Student)
					if !ok {
						panic("传入参数不是model.Student变量")
					}
					return Student.ID, nil
				}},
			"created_at": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					Student, ok := (p.Source).(model.Student)
					if !ok {
						panic("传入参数不是model.Student变量")
					}
					return Student.CreatedAt, nil
				},
			},
			"updated_at": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					Student, ok := (p.Source).(model.Student)
					if !ok {
						panic("传入参数不是model.Student变量")
					}
					return Student.UpdatedAt, nil
				},
			},
			"deleted_at": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					Student, ok := (p.Source).(model.Student)
					if !ok {
						panic("传入参数不是model.Student变量")
					}
					return Student.DeletedAt, nil
				},
			},
			"name": &graphql.Field{Type: graphql.String},
			"key":  &graphql.Field{Type: graphql.String},
		},
		Description: "学生",
	},
)

func StudentList() *graphql.Field {
	return &graphql.Field{
		Description: "学生列表",
		Type:        graphql.NewList(&StudentObject),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result []model.Student
			id, ok := p.Args["id"].(int)
			db := mysql.GetIns().Model(&result)
			if ok {
				db.Where("`id` = ?", id)
			}
			db.Find(&result)
			return result, nil
		},
	}
}

func StudentShow() *graphql.Field {
	return &graphql.Field{
		Description: "学生",
		Type:        &StudentObject,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.Student
			id, ok := p.Args["id"].(int)
			db := mysql.GetIns().Model(&result)
			if ok {
				db.Where("`id` = ?", id)
			}
			db.First(&result)
			return result, nil
		},
	}
}
