package fitness_test

import "github.com/satori/go.uuid"

type Question struct {
	UUID 	uuid.UUID
	Text     string `json:"text" bson:"text"`
	Answers  []Answer `json:"answers" bson:"answers"`
	Weight   float64 `json:"weight" bson:"weight"`
	Category Category
}


type Answer struct {
	UUID   uuid.UUID
	Text  string `json:"text" bson:"text"`
	Score float64 `json:"score" bson:"score"`
}

type Category struct {
	Name        string`json:"name" bson:"name"`
	Description string`json:"description" bson:"description"`
}
