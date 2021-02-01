package field

import (
	"github.com/graphql-go/graphql"
)

func getConnectionFunc(f func() *graphql.Field, name string, description string) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:        name,
			Description: description,
			Fields: graphql.Fields{
				"totalCount": &graphql.Field{
					Type: graphql.Int,
					//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					//	fmt.Println("totalCount 开始")
					//	//fmt.Println("Info",p.Info)
					//	//fmt.Println("Source",p.Source)
					//	//fmt.Println("Context",p.Context)
					//	//fmt.Println("Args",p.Args)
					//	//未完成
					//	return 1, nil
					//},
				},
				"edges": &graphql.Field{
					Type: graphql.NewObject(
						graphql.ObjectConfig{
							Name:        "Edges",
							Description: "列表",
							Fields: graphql.Fields{
								"node": f(),
								"cursor": &graphql.Field{
									Type: graphql.String,
									//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
									//	fmt.Println("cursor 开始")
									//	//fmt.Println("Info",p.Info)
									//	//fmt.Println("Source",p.Source)
									//	//fmt.Println("Context",p.Context)
									//	//fmt.Println("Args",p.Args)
									//	//未完成
									//	return 1, nil
									//},
								},
							},
						},
					),
					//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					//	fmt.Println("edges 开始")
					//	//未完成
					//	return p.Source, nil
					//},
				},
				"pageInfo": &graphql.Field{
					Type: graphql.NewObject(
						graphql.ObjectConfig{
							Name:        "PageInfo",
							Description: "分页信息",
							Fields: graphql.Fields{
								"hasNextPage": &graphql.Field{
									Type: graphql.Boolean,
									//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
									//	fmt.Println("hasNextPage 开始")
									//	//fmt.Println("Info",p.Info)
									//	//fmt.Println("Source",p.Source)
									//	//fmt.Println("Context",p.Context)
									//	//fmt.Println("Args",p.Args)
									//	//未完成
									//	return 1, nil
									//},
								},
								"endCursor": &graphql.Field{
									Type: graphql.String,
									//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
									//	fmt.Println("endCursor 开始")
									//	//fmt.Println("Info",p.Info)
									//	//fmt.Println("Source",p.Source)
									//	//fmt.Println("Context",p.Context)
									//	//fmt.Println("Args",p.Args)
									//	//未完成
									//	return 1, nil
									//},
								},
							},
						},
					),
					//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					//	fmt.Println("pageInfo 开始")
					//	//fmt.Println("Info",p.Info)
					//	//fmt.Println("Source",p.Source)
					//	//fmt.Println("Context",p.Context)
					//	//fmt.Println("Args",p.Args)
					//	var result = make(map[string]interface{})
					//	result["endCursor"] = "3"
					//	result["hasNextPage"] = true
					//	//未完成
					//	return result, nil
					//},
				},
			},
		},
	)
}

func getConnectionList(list *graphql.List, name string, description string) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:        name,
			Description: description,
			Fields: graphql.Fields{
				"totalCount": &graphql.Field{
					Description: "字段:总数",
					Type:        graphql.Int,
					//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					//	fmt.Println("totalCount 开始")
					//	//fmt.Println("Info",p.Info)
					//	//fmt.Println("Source",p.Source)
					//	//fmt.Println("Context",p.Context)
					//	//fmt.Println("Args",p.Args)
					//	//未完成
					//	return 1, nil
					//},
				},
				"edges": &graphql.Field{
					Description: "字段：列表信息",
					Type: graphql.NewObject(
						graphql.ObjectConfig{
							Name:        "Edges",
							Description: "对象:列表信息",
							Fields: graphql.Fields{
								"node": &graphql.Field{
									Description: "节点",
									Type:        list,
								},
								"cursor": &graphql.Field{
									Description: "第一个节点游标",
									Type:        graphql.String,
									//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
									//	fmt.Println("cursor 开始")
									//	//fmt.Println("Info",p.Info)
									//	//fmt.Println("Source",p.Source)
									//	//fmt.Println("Context",p.Context)
									//	//fmt.Println("Args",p.Args)
									//	//未完成
									//	return 1, nil
									//},
								},
							},
						},
					),
					//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					//	fmt.Println("edges 开始")
					//	//未完成
					//	return p.Source, nil
					//},
				},
				"pageInfo": &graphql.Field{
					Description: "字段：分页信息",
					Type: graphql.NewObject(
						graphql.ObjectConfig{
							Name:        "PageInfo",
							Description: "对象:分页信息",
							Fields: graphql.Fields{
								"hasNextPage": &graphql.Field{
									Description: "字段：是否有下一页",
									Type:        graphql.Boolean,
									//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
									//	fmt.Println("hasNextPage 开始")
									//	//fmt.Println("Info",p.Info)
									//	//fmt.Println("Source",p.Source)
									//	//fmt.Println("Context",p.Context)
									//	//fmt.Println("Args",p.Args)
									//	//未完成
									//	return 1, nil
									//},
								},
								"endCursor": &graphql.Field{
									Description: "字段：最后一个节点游标",
									Type:        graphql.String,
									//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
									//	fmt.Println("endCursor 开始")
									//	//fmt.Println("Info",p.Info)
									//	//fmt.Println("Source",p.Source)
									//	//fmt.Println("Context",p.Context)
									//	//fmt.Println("Args",p.Args)
									//	//未完成
									//	return 1, nil
									//},
								},
								"StartCursor": &graphql.Field{
									Description: "字段：第一个节点游标",
									Type:        graphql.String,
									//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
									//	fmt.Println("endCursor 开始")
									//	//fmt.Println("Info",p.Info)
									//	//fmt.Println("Source",p.Source)
									//	//fmt.Println("Context",p.Context)
									//	//fmt.Println("Args",p.Args)
									//	//未完成
									//	return 1, nil
									//},
								},
							},
						},
					),
					//Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					//	fmt.Println("pageInfo 开始")
					//	//fmt.Println("Info",p.Info)
					//	//fmt.Println("Source",p.Source)
					//	//fmt.Println("Context",p.Context)
					//	//fmt.Println("Args",p.Args)
					//	var result = make(map[string]interface{})
					//	result["endCursor"] = "3"
					//	result["hasNextPage"] = true
					//	//未完成
					//	return result, nil
					//},
				},
			},
		},
	)
}
