package api

import "github.com/graphql-go/graphql"

var mutationFields = graphql.Fields{
	"saveTest" : &graphql.Field {
		Type:graphql.String,
		Args: graphql.FieldConfigArgument{
			"userId": &graphql.ArgumentConfig{
				Type: graphql.String,
			},

		},
		Resolve:func(p graphql.ResolveParams) (interface{}, error){
			return "test saved", nil
		},


	},
}
