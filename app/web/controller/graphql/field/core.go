package field

import "github.com/graphql-go/graphql"

func getConnection(f func() *graphql.Field) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Fields: graphql.Fields{
				"totalCount": &graphql.Field{Type: graphql.Int},
				"edges": &graphql.Field{
					Type: graphql.NewObject(
						graphql.ObjectConfig{
							Fields: graphql.Fields{
								"node":   f(),
								"cursor": &graphql.Field{Type: graphql.String},
							},
						},
					),
				},
				"pageInfo": &graphql.Field{
					Type: graphql.NewObject(
						graphql.ObjectConfig{
							Fields: graphql.Fields{
								"hasNextPage": &graphql.Field{Type: graphql.Boolean},
								"endCursor":   &graphql.Field{Type: graphql.String},
							},
						},
					),
				},
			},
		},
	)
}
