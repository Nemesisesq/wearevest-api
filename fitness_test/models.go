package fitness_test

import (
	"github.com/satori/go.uuid"
	"github.com/fatih/structs"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"encoding/json"
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
	UUID   uuid.UUID `json:"uuid" bson:"uuid"`
	Q      Question`json:"q" bson:"q"`
	A      Answer`json:"a" bson:"a"`
	Result Result `json:"result" bson:"result"`
}

func (i *Interchange) prepare(db *mgo.Database) map[string]interface{} {

	return structs.Map(i)
}

func (i *Interchange) collection() string {

	return "interchanges"
}

func (i *Interchange) getUUID() string {

	return i.UUID.String()

}

type Result struct {
	RawScore      float64`json:"raw_score" bson:"raw_score"`
	WeightedScore float64`json:"weighted_score" bson:"weighted_score"`
}

/*
Updater type
*/
type Updater interface {
	prepare(db *mgo.Database) map[string]interface{}
	collection() string
	getUUID() string
}

/*
Updater Methods
*/
func Update(updater Updater, db *mgo.Database) {
	db.C(updater.collection()).Upsert(bson.M{"uuid": updater.getUUID()}, updater.prepare(db))
}

type Inflater interface {
}
