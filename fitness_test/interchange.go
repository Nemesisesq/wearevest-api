package fitness_test

import (
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
)

type Interchange struct {
	UUID   uuid.UUID `json:"uuid" bson:"uuid"`
	Q      Question`json:"q" bson:"q"`
	A      Answer`json:"a" bson:"a"`
	Result Result `json:"result" bson:"result"`
}

type interchange struct {
	UUID uuid.UUID `json:"uuid" bson:"uuid"`
	Q      uuid.UUID`json:"q" bson:"q"`
	A      Answer`json:"a" bson:"a"`
	Result Result `json:"result" bson:"result"`
}

func (i Interchange) prepare(db *mgo.Database) interface{} {

	prepped := &interchange{i.UUID, i.Q.UUID, i.A, i.Result}
	return *prepped
}

func (i Interchange) collection() string {

	return "interchanges"
}

func (i Interchange) getUUID() string {

	return i.UUID.String()

}
