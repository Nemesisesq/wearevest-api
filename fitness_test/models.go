package fitness_test

import (
	"github.com/satori/go.uuid"
	"github.com/compose/transporter/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Answer struct {
	UID   uuid.UUID
	Text  string `json:"text" bson:"text"`
	Score float64 `json:"score" bson:"score"`
}

type Category struct {
	Name        string`json:"name" bson:"name"`
	Description string`json:"description" bson:"description"`
}

type Question struct {
	Text     string `json:"text" bson:"text"`
	Answers  []Answer `json:"answers" bson:"answers"`
	Weight   float64 `json:"weight" bson:"weight"`
	Category Category
}

type User struct {
	UUID      uuid.UUID`json:"uuid" bson:"uuid"`
	FirstName string`json:"first_name" bson:"first_name"`
	LastName  string`json:"last_name" bson:"last_name"`
	Email     string`json:"email" bson:"email"`
}

type Interchange struct {
	Q      Question`json:"q" bson:"q"`
	A      Answer`json:"a" bson:"a"`
	Result Result `json:"result" bson:"result"`
}

type Result struct {
	RawScore      float64`json:"raw_score" bson:"raw_score"`
	WeightedScore float64`json:"weighted_score" bson:"weighted_score"`
}

type Test struct {
	UUID         uuid.UUID`json:"uuid" bson:"uuid"`
	Interchanges []Interchange`json:"interchanges" bson:"interchanges"`
	User         User `json:"user" bson:"user"`
	Result       Result`json:"result" bson:"result"`
}

func (t *Test) prepare(db *mgo.Database) ( m map[string]interface{}) {


	for _, v := range t.Interchanges {
		Update(v, db)
	}


	return m
}

func (t *Test) collection() string {
	return "tests"
}

func (t *Test) GetUUID() string {
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

type Updater interface {
	prepare() map[string]interface{}
	collection() string
	getUUID() string
}

func Update(updater Updater, db *mgo.Database) {
	db.C(updater.collection()).Upsert(bson.M{"uuid": updater.getUUID()}, updater.prepare() )
}

type Inflater interface {
}
