package fitness_test

import (

	"context"
	"gopkg.in/mgo.v2"
	"github.com/nemesisesq/wearevest/fitness_test/models"
)

func GetQuestions(context context.Context) []map[string]interface{} {


	questions := []map[string]interface{}{}
	db := context.Value("db").(*mgo.Database)

	col := db.C("questions")

	col.Find(nil).All(questions)

	return questions



}

func ComputeFitnessTestResults(test fitness_test.Test) {
	for _, v := range test.Interchanges {
		v.Result.RawScore = v.A.Score
		v.Result.WeightedScore = v.A.Score * v.Q.Weight
	}

	test.ComputeTotalScores()

}
