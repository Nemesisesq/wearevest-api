package fitness_test

import (
	"context"
	"reflect"
	"testing"
	"github.com/nemesisesq/wearevest-api/shared/models"
	"gopkg.in/mgo.v2"
	"github.com/satori/go.uuid"
)

func TestGetQuestions(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		args args
		want []map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQuestions(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQuestions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeFitnessTestResults(t *testing.T) {
	type args struct {
		test Test
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ComputeFitnessTestResults(tt.args.test)
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		updater Updater
		db      *mgo.Database
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Update(tt.args.updater, tt.args.db)
		})
	}
}

func TestInflate(t *testing.T) {
	type args struct {
		inflater Inflater
		db       *mgo.Database
		uuid     uuid.UUID
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Inflate(tt.args.inflater, tt.args.db, tt.args.uuid)
		})
	}
}


func TestInterchange_prepare(t *testing.T) {
	type fields struct {
		UUID   uuid.UUID
		Q      Question
		A      Answer
		Result Result
	}
	type args struct {
		db *mgo.Database
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interchange{
				UUID:   tt.fields.UUID,
				Q:      tt.fields.Q,
				A:      tt.fields.A,
				Result: tt.fields.Result,
			}
			if got := i.prepare(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interchange.prepare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterchange_collection(t *testing.T) {
	type fields struct {
		UUID   uuid.UUID
		Q      Question
		A      Answer
		Result Result
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interchange{
				UUID:   tt.fields.UUID,
				Q:      tt.fields.Q,
				A:      tt.fields.A,
				Result: tt.fields.Result,
			}
			if got := i.collection(); got != tt.want {
				t.Errorf("Interchange.collection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterchange_getUUID(t *testing.T) {
	type fields struct {
		UUID   uuid.UUID
		Q      Question
		A      Answer
		Result Result
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interchange{
				UUID:   tt.fields.UUID,
				Q:      tt.fields.Q,
				A:      tt.fields.A,
				Result: tt.fields.Result,
			}
			if got := i.getUUID(); got != tt.want {
				t.Errorf("Interchange.getUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTest_prepare(t *testing.T) {
	type fields struct {
		UUID         uuid.UUID
		Interchanges []Interchange
		User         shared.User
		Result       Result
	}
	type args struct {
		db *mgo.Database
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &Test{
				UUID:         tt.fields.UUID,
				Interchanges: tt.fields.Interchanges,
				User:         tt.fields.User,
				Result:       tt.fields.Result,
			}
			if got := t.prepare(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test.prepare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTest_collection(t *testing.T) {
	type fields struct {
		UUID         uuid.UUID
		Interchanges []Interchange
		User         shared.User
		Result       Result
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := Test{
				UUID:         tt.fields.UUID,
				Interchanges: tt.fields.Interchanges,
				User:         tt.fields.User,
				Result:       tt.fields.Result,
			}
			if got := t.collection(); got != tt.want {
				t.Errorf("Test.collection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTest_getUUID(t *testing.T) {
	type fields struct {
		UUID         uuid.UUID
		Interchanges []Interchange
		User         shared.User
		Result       Result
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := Test{
				UUID:         tt.fields.UUID,
				Interchanges: tt.fields.Interchanges,
				User:         tt.fields.User,
				Result:       tt.fields.Result,
			}
			if got := t.getUUID(); got != tt.want {
				t.Errorf("Test.getUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTest_ComputeTotalScores(t *testing.T) {
	type fields struct {
		UUID         uuid.UUID
		Interchanges []Interchange
		User         shared.User
		Result       Result
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := Test{
				UUID:         tt.fields.UUID,
				Interchanges: tt.fields.Interchanges,
				User:         tt.fields.User,
				Result:       tt.fields.Result,
			}
			t.ComputeTotalScores()
		})
	}
}
