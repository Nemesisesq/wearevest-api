package fitness_test

import (
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"github.com/nemesisesq/wearevest/shared/models"
	"github.com/nemesisesq/wearevest/fitness_test"
)

type Test struct {
	UUID         uuid.UUID`json:"uuid" bson:"uuid"`
	Interchanges []Interchange`json:"interchanges" bson:"interchanges"`
	User         shared.User `json:"user" bson:"user"`
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

/*
updated methods
*/
func (t *Test) prepare(db *mgo.Database) interface{} {


	interchanges := []uuid.UUID{}
	for _, v := range t.Interchanges {
		fitness_test.Update(v, db)
		interchanges = append(interchanges, v.UUID)
	}

	prepped := &test{UUID:t.UUID, User:t.User.UUID, Result:t.Result}
	prepped.Interchanges = interchanges

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
