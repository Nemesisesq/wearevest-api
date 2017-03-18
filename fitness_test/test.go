package fitness_test

import (
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
)

type Test struct {
	UUID         uuid.UUID`json:"uuid" bson:"uuid"`
	Interchanges []Interchange`json:"interchanges" bson:"interchanges"`
	User         User `json:"user" bson:"user"`
	Result       Result`json:"result" bson:"result"`
}

/*
This struct is for saving and retrieving from the database
*/

type test struct {
	UUID         uuid.UUID`json:"uuid" bson:"uuid"`
	Interchanges []uuid.UUID`json:"interchanges" bson:"interchanges"`
	User         uuid.UUID `json:"user" bson:"user"`
	Result       Result `json:"result" bson:"result"`
}

func (t *Test) prepare(db *mgo.Database) test {

	interchanges := []uuid.UUID{}
	for _, v := range t.Interchanges {
		Update(v, db)
		interchanges = append(interchanges, v.UUID)
	}

	prepped := &test{UUID:t.UUID, User:t.User.UUID, Result:t.Result}

	return *prepped
}

func (t *Test) collection() string {
	return "tests"
}

func (t *Test) getUUID() string {
	return t.UUID.String()
}

func (t Test) ComputeTotalScores() {
	for _, v := range t.Interchanges {
		t.Result.RawScore += v.Result.RawScore
	}

	for _, v := range t.Interchanges {
		t.Result.WeightedScore += v.Result.WeightedScore
	}
}
