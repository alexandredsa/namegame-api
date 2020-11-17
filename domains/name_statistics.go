package domains

type NameStatistics struct {
	Name  string `bson:"name"`
	Total int32  `bson:"total"`
}
