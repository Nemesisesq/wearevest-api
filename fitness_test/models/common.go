package fitness_test


type Result struct {
	RawScore      float64`json:"raw_score" bson:"raw_score"`
	WeightedScore float64`json:"weighted_score" bson:"weighted_score"`
}
