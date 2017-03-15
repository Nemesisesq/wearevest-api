package api

import (
	"net/http"
	"github.com/graphql-go/graphql"
	"encoding/json"
	log"github.com/Sirupsen/logrus"
)

var Schema graphql.Schema

var fields graphql.Fields

var queryType *graphql.Object

func init() {

	fields = graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	queryType = graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Fields: fields})

	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
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
	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: r.URL.Query()["query"][0],
		//Context:      r.WithContext(context.Background(), "currentUser", user),
	})
	if len(result.Errors) > 0 {
		log.Printf("wrong result, unexpected errors: %v", result.Errors)
		return
	}
	json.NewEncoder(w).Encode(result)
}
