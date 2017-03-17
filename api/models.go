package api

import "github.com/graphql-go/graphql"

var answerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Answer",
		Fields: graphql.Fields{
			"text" : &graphql.Field{
				Type: graphql.String,
			},
		},
})

var answersType = graphql.NewList(answerType)

var questionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Question",
		Fields: graphql.Fields{
			"text" : &graphql.Field{
				Type: graphql.String,
			},
			"answers" : &graphql.Field{
				Type: answersType,
			},
		},
})

var questionsType = graphql.NewList(questionType)
