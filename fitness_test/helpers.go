package fitness_test

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/satori/go.uuid"
)




/*
Updater type
*/
type Updater interface {
	prepare(db *mgo.Database) interface{}
	collection() string
	getUUID() string
}

/*
Updater Methods
*/
func Update(updater Updater, db *mgo.Database) {
	db.C(updater.collection()).Upsert(bson.M{"uuid": updater.getUUID()}, updater.prepare(db))
}

/*
Inflater type
*/


type Inflater interface {
	repair(db *mgo.Database) interface{}
	collection() string
	getUUID() string
}

func Inflate(inflater Inflater, db *mgo.Database, uuid uuid.UUID)  {

	db.C(inflater.collection()).Find(bson.M{"uuid": uuid}).One(&inflater)

}
