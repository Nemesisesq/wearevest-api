package api

import (
	"github.com/graphql-go/graphql"
	"github.com/nemesisesq/wearevest-api/fitness_test"
)

var queryFields = graphql.Fields{
	"hello": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	},

	"questions": &graphql.Field{
		Type: questionsType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return fitness_test.GetQuestions(p.Context), nil
		},
	},

}
