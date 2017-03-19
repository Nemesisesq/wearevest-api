package api

import (
	"net/http"
	"github.com/graphql-go/graphql"
	"encoding/json"
	log"github.com/Sirupsen/logrus"
)

var Schema graphql.Schema



var queryType *graphql.Object

var mutationType *graphql.Object

func init() {

	queryType = graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields})
	mutationType = graphql.NewObject(graphql.ObjectConfig{Name: "RootMutation", Fields: mutationFields })

	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
		Mutation: mutationType,
	})
	if err != nil {
		log.Fatalf("failed to create schema, error: %v", err)
	}

	Schema = s
}

func GraphqlHandler(w http.ResponseWriter, r *http.Request) {
	//user := struct {
	//	ID   int    `json:"id"`
	//	Name string `json:"name"`
	//}{1, "cool user"}

	query :=r.URL.Query()["query"][0]
	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: query,
		Context:      r.Context(),
	})
	if len(result.Errors) > 0 {
		log.Printf("wrong result, unexpected errors: %v", result.Errors)
		return
	}
	json.NewEncoder(w).Encode(result)
}
