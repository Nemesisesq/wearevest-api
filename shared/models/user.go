package shared

import "github.com/satori/go.uuid"

type User struct {
	UUID      uuid.UUID`json:"uuid" bson:"uuid"`
	FirstName string`json:"first_name" bson:"first_name"`
	LastName  string`json:"last_name" bson:"last_name"`
	Email     string`json:"email" bson:"email"`
}
