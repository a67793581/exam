package field

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

var StudentObject = *graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Student",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
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
			var where = make(map[string]interface{})
			db := mysql.GetIns().Model(&result)
			id, ok := p.Args["id"].(int)
			if ok {
				where["id"] = id
			}
			t := reflect.TypeOf(p.Source)
			v := reflect.ValueOf(p.Source)
			if t.Kind() == reflect.Struct {
				StructField, ok := t.FieldByName("StudentID")
				if ok {
					where["id"] = v.FieldByName(StructField.Name).Interface()
				}
			}
			fmt.Println("条件", where)
			db.Where(where)
			db.First(&result)
			return result, nil
		},
	}
}

func StudentConnection() *graphql.Field {
	return &graphql.Field{
		Description: "字段:学生列表信息",
		Type:        getConnectionList(graphql.NewList(&StudentObject), "StudentObjectList", "对象：学生列表信息"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"key": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"first": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: 5,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"after": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"before": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var list []model.Student
			db := mysql.GetIns()
			mDb := db.Model(&list)
			var where = make(map[string]interface{})
			id, ok := p.Args["id"].(int)
			if ok {
				where["id"] = id
			}
			key, ok := p.Args["key"].(string)
			if ok {
				where["key"] = key
			}
			var totalCount int64
			var afterCount int64
			mDb.Where(where)
			mDb.Count(&totalCount)
			after, ok := p.Args["after"].(int)
			if ok {
				mDb.Where("`id` > ?", after)
			}
			isDesc := false
			before, ok := p.Args["before"].(int)
			if ok {
				mDb.Where("`id` < ?", before)
				mDb.Order("`id` desc")
				isDesc = true
			}

			first, ok := p.Args["first"].(int)
			if ok {
				mDb.Limit(first)
			}
			offset, ok := p.Args["offset"].(int)
			if ok {
				mDb.Offset(offset)
			}
			mDb.Find(&list)
			if isDesc {
				length := len(list)
				for i := 0; i < length/2; i++ {
					temp := list[length-1-i]
					list[length-1-i] = list[i]
					list[i] = temp
				}
			}

			if len(list) > 0 {
				db.Model(&list).Where("`id` > ?", list[len(list)-1].ID).Count(&afterCount)
			}

			pageInfo := make(map[string]interface{})
			pageInfo["first"] = first
			pageInfo["totalCount"] = totalCount
			pageInfo["afterCount"] = afterCount
			if len(list) > 0 {
				if afterCount > 0 {
					pageInfo["hasNextPage"] = afterCount > 0
				} else {
					pageInfo["hasNextPage"] = (float64(totalCount) / float64(first)) > 1
				}
				pageInfo["endCursor"] = list[len(list)-1].ID
				pageInfo["startCursor"] = list[0].ID
			} else {
				pageInfo["hasNextPage"] = false
				pageInfo["startCursor"] = 0
				pageInfo["endCursor"] = 0
			}

			var result = make(map[string]interface{})
			result["list"] = list
			result["pageInfo"] = pageInfo
			return result, nil
		},
	}
}

func StudentDML() *graphql.Field {
	return &graphql.Field{
		Description: "字段:学生-增删改",
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "StudentDML",
			Description: "对象:学生-增删改",
			Fields: graphql.Fields{
				"create": StudentCreate(),
				"update": StudentUpdate(),
				"delete": StudentDelete(),
			},
		}),

		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return p.Source, nil
		},
	}
}

func StudentCreate() *graphql.Field {
	return &graphql.Field{
		Description: "创建学生",
		Type:        &StudentObject,
		Args: graphql.FieldConfigArgument{
			"key": &graphql.ArgumentConfig{
				Description: "学号",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.ArgumentConfig{
				Description: "姓名",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.Student
			db := mysql.GetIns().Model(&result)
			Name, _ := p.Args["name"].(string)
			result.Name = Name
			if Name == "" {
				panic("姓名 不允许为空")
			}
			Key, _ := p.Args["key"].(string)
			result.Key = Key
			if Key == "" {
				panic("学号 不允许为空")
			}
			db.Create(&result)
			return result, nil
		},
	}
}

func StudentUpdate() *graphql.Field {
	return &graphql.Field{
		Description: "更新学生",
		Type:        &StudentObject,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "自增ID",
				Type:        graphql.NewNonNull(graphql.Int),
			},
			"key": &graphql.ArgumentConfig{
				Description: "学号",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.ArgumentConfig{
				Description: "姓名",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.Student
			db := mysql.GetIns().Model(&result)
			ID, _ := p.Args["id"].(int)
			result.ID = uint(ID)
			Error := db.First(&result).Error
			// 检查错误
			if Error != nil {
				panic(Error)
			}

			Name, _ := p.Args["name"].(string)
			result.Name = Name
			if Name == "" {
				panic("姓名 不允许为空")
			}
			Key, _ := p.Args["key"].(string)
			result.Key = Key
			if Key == "" {
				panic("学号 不允许为空")
			}
			db.Save(&result)
			return result, nil
		},
	}
}

func StudentDelete() *graphql.Field {
	return &graphql.Field{
		Description: "删除学生",
		Type:        &StudentObject,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "自增ID",
				Type:        graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.Student
			db := mysql.GetIns().Model(&result)
			ID, _ := p.Args["id"].(int)
			result.ID = uint(ID)
			Error := db.First(&result).Error
			// 检查错误
			if Error != nil {
				panic(Error)
			}
			db.Delete(&result)
			return result, nil
		},
	}
}
