package field

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"fmt"
	"github.com/graphql-go/graphql"
)

type ExamRecord struct {
	Object *graphql.Object
}

var ExamRecordObject = *graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "ExamRecord",
		Description: "考试记录",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.ID, nil
					}
					return 0, nil
				}},
			"created_at": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.CreatedAt, nil
					}
					return 0, nil
				},
			},
			"updated_at": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.UpdatedAt, nil
					}
					return 0, nil
				},
			},
			"deleted_at": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.DeletedAt, nil
					}
					return 0, nil
				},
			},
			"code": &graphql.Field{
				Type:        graphql.String,
				Description: "考场批次",
			},
			"key": &graphql.Field{
				Type:        graphql.String,
				Description: "考试唯一编码",
			},
			"exam_time": &graphql.Field{
				Type:        graphql.Int,
				Description: "考试时间",
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.ExamTime, nil
					}
					return 0, nil
				},
			},
			"achievement": &graphql.Field{
				Type:        graphql.Int,
				Description: "考试成绩",
			},
			"course_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "课程id",
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.CourseID, nil
					}
					return 0, nil
				},
			},
			"student_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "学生id",
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					ExamRecord, ok := (p.Source).(model.ExamRecord)
					if ok {
						return ExamRecord.StudentID, nil
					}
					return 0, nil
				},
			},
			"student": StudentShow(),
			"course":  CourseShow(),
		},
	},
)

func ExamRecordList() *graphql.Field {
	return &graphql.Field{
		Description: "考试记录列表",
		Type:        graphql.NewList(&ExamRecordObject),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"key": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"first": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"after": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			var examRecords []model.ExamRecord
			db := mysql.GetIns().Model(&examRecords)
			var where = make(map[string]interface{})
			id, ok := p.Args["id"].(int)
			if ok {
				where["id"] = id
			}
			key, ok := p.Args["key"].(string)
			if ok {
				where["key"] = key
			}
			first, ok := p.Args["first"].(int)
			if ok {
				db.Limit(first)
			}
			offset, ok := p.Args["offset"].(int)
			if ok {
				db.Offset(offset)
			}
			after, ok := p.Args["after"].(int)
			if ok {
				db.Where("`id` > '?'", after)
			}
			db.Where(where)

			db.Find(&examRecords)
			//var count int64
			//db.Count(&count)
			//var result = make(map[string]interface{})
			//result["edges"] = examRecords
			//result["totalCount"] = count
			//pageInfo := make(map[string]interface{})
			//if len(examRecords) > 0 {
			//	pageInfo["startCursor"] = examRecords[0]
			//	pageInfo["endCursor"] = examRecords[len(examRecords)-1]
			//}
			//result["pageInfo"] = pageInfo
			return examRecords, nil
		},
	}
}

func ExamRecordShow() *graphql.Field {
	return &graphql.Field{
		Description: "考试记录",
		Type:        &ExamRecordObject,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"key": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.ExamRecord
			db := mysql.GetIns().Model(&result)
			var where = make(map[string]interface{})
			id, ok := p.Args["id"].(int)
			if ok {
				where["id"] = id
			}
			key, ok := p.Args["key"].(string)
			if ok {
				where["key"] = key
			}
			db.Where(where)
			db.First(&result)
			return result, nil
		},
	}
}

func ExamRecordConnection() *graphql.Field {
	return &graphql.Field{
		Description: "字段:考试记录列表信息",
		Type:        getConnectionList(graphql.NewList(&ExamRecordObject), "ExamRecordObjectList", "对象：考试记录列表信息"),
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
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var examRecords []model.ExamRecord
			db := mysql.GetIns().Model(&examRecords)
			var where = make(map[string]interface{})
			id, ok := p.Args["id"].(int)
			if ok {
				where["id"] = id
			}
			key, ok := p.Args["key"].(string)
			if ok {
				where["key"] = key
			}
			first, ok := p.Args["first"].(int)
			if ok {
				db.Limit(first)
			}
			offset, ok := p.Args["offset"].(int)
			if ok {
				db.Offset(offset)
			}
			db.Where(where)
			var totalCount int64
			var afterCount int64
			db.Count(&totalCount)
			after, ok := p.Args["after"].(int)
			if ok {
				db.Where("`id` > '?'", after)
				db.Count(&afterCount)
			}
			db.Find(&examRecords)
			edges := make(map[string]interface{})
			edges["node"] = examRecords
			pageInfo := make(map[string]interface{})
			if afterCount > 0 {
				pageInfo["hasNextPage"] = afterCount > 0
			} else {
				pageInfo["hasNextPage"] = (float64(totalCount) / float64(first)) > 1
			}
			if len(examRecords) > 0 {
				edges["cursor"] = examRecords[0].ID
				pageInfo["endCursor"] = examRecords[len(examRecords)-1].ID
			} else {
				edges["cursor"] = 0
				pageInfo["endCursor"] = 0
			}

			var result = make(map[string]interface{})
			result["totalCount"] = totalCount
			result["edges"] = edges
			result["pageInfo"] = pageInfo
			return result, nil
		},
	}
}

func ExamRecordDML() *graphql.Field {
	return &graphql.Field{
		Description: "字段:考试记录-增删改",
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "ExamRecordDML",
			Description: "对象:考试记录-增删改",
			Fields: graphql.Fields{
				"create": ExamRecordCreate(),
				"update": ExamRecordUpdate(),
				"delete": ExamRecordDelete(),
			},
		}),

		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return p.Source, nil
		},
	}
}

func ExamRecordCreate() *graphql.Field {
	return &graphql.Field{
		Description: "创建考试记录",
		Type:        &ExamRecordObject,
		Args: graphql.FieldConfigArgument{
			"key": &graphql.ArgumentConfig{
				Description: "考试唯一编码",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"code": &graphql.ArgumentConfig{
				Description: "考场批次",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"exam_time": &graphql.ArgumentConfig{
				Description: "考场批次",
				Type:        graphql.NewNonNull(graphql.Int),
			},
			"achievement": &graphql.ArgumentConfig{
				Description: "考试成绩",
				Type:        graphql.NewNonNull(graphql.Int),
			},
			"course_id": &graphql.ArgumentConfig{
				Description: "课程id",
				Type:        graphql.NewNonNull(graphql.Int),
			},
			"student_id": &graphql.ArgumentConfig{
				Description: "学生id",
				Type:        graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.ExamRecord
			db := mysql.GetIns().Model(&result)
			StudentID, _ := p.Args["student_id"].(int32)
			result.StudentID = StudentID
			CourseID, _ := p.Args["course_id"].(int32)
			result.CourseID = CourseID
			Achievement, _ := p.Args["achievement"].(int32)
			result.Achievement = Achievement
			ExamTime, _ := p.Args["exam_time"].(int32)
			result.ExamTime = ExamTime
			Code, _ := p.Args["code"].(string)
			result.Code = Code
			Key, _ := p.Args["key"].(string)
			result.Key = Key
			db.Create(&result)
			fmt.Println(result)
			return result, nil
		},
	}
}

func ExamRecordUpdate() *graphql.Field {
	return &graphql.Field{
		Description: "更新考试记录",
		Type:        &ExamRecordObject,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "自增ID",
				Type:        graphql.NewNonNull(graphql.Int),
			},
			"key": &graphql.ArgumentConfig{
				Description: "考试唯一编码",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"code": &graphql.ArgumentConfig{
				Description: "考场批次",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"exam_time": &graphql.ArgumentConfig{
				Description: "考场批次",
				Type:        graphql.NewNonNull(graphql.Int),
			},
			"achievement": &graphql.ArgumentConfig{
				Description: "考试成绩",
				Type:        graphql.NewNonNull(graphql.Int),
			},
			"course_id": &graphql.ArgumentConfig{
				Description: "课程id",
				Type:        graphql.NewNonNull(graphql.Int),
			},
			"student_id": &graphql.ArgumentConfig{
				Description: "学生id",
				Type:        graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.ExamRecord
			db := mysql.GetIns().Model(&result)
			ID, _ := p.Args["id"].(int)
			result.ID = uint(ID)
			Error := db.First(&result).Error
			// 检查错误
			if Error != nil {
				panic(Error)
			}

			StudentID, _ := p.Args["student_id"].(int32)
			result.StudentID = StudentID
			CourseID, _ := p.Args["course_id"].(int32)
			result.CourseID = CourseID
			Achievement, _ := p.Args["achievement"].(int32)
			result.Achievement = Achievement
			ExamTime, _ := p.Args["exam_time"].(int32)
			result.ExamTime = ExamTime
			Code, _ := p.Args["code"].(string)
			result.Code = Code
			Key, _ := p.Args["key"].(string)
			result.Key = Key
			db.Save(&result)
			fmt.Println(result)
			return result, nil
		},
	}
}

func ExamRecordDelete() *graphql.Field {
	return &graphql.Field{
		Description: "删除考试记录",
		Type:        &ExamRecordObject,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "自增ID",
				Type:        graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var result model.ExamRecord
			db := mysql.GetIns().Model(&result)
			ID, _ := p.Args["id"].(int)
			result.ID = uint(ID)
			Error := db.First(&result).Error
			// 检查错误
			if Error != nil {
				panic(Error)
			}
			db.Delete(&result)
			fmt.Println(result)
			return result, nil
		},
	}
}
