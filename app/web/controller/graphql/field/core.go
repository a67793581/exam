package field

import (
	"github.com/graphql-go/graphql"
)

var PageInfo = *graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "PageInfo",
		Description: "对象:分页信息",
		Fields: graphql.Fields{
			"hasNextPage": &graphql.Field{
				Description: "字段：是否有下一页",
				Type:        graphql.Boolean,
			},
			"endCursor": &graphql.Field{
				Description: "字段：最后一个节点游标",
				Type:        graphql.Int,
			},
			"startCursor": &graphql.Field{
				Description: "字段：第一个节点游标",
				Type:        graphql.Int,
			},
			"first": &graphql.Field{
				Description: "字段：每页条数",
				Type:        graphql.Int,
			},
			"totalCount": &graphql.Field{
				Description: "字段:总数",
				Type:        graphql.Int,
			},
			"afterCount": &graphql.Field{
				Description: "字段:剩余总数",
				Type:        graphql.Int,
			},
		},
	},
)

func getConnectionList(list *graphql.List, name string, description string) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:        name,
			Description: description,
			Fields: graphql.Fields{
				"pageInfo": &graphql.Field{
					Description: "字段：分页信息",
					Type:        &PageInfo,
				},
				"list": &graphql.Field{
					Description: "列表",
					Type:        list,
				},
			},
		},
	)
}
