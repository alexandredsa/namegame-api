package domains

type NameStatistics struct {
	Name      string `bson:"Nome"`
	Until1930 int    `bson:"ate1930"`
	Until1940 int    `bson:"ate1940"`
	Until1950 int    `bson:"ate1950"`
	Until1960 int    `bson:"ate1960"`
	Until1970 int    `bson:"ate1970"`
	Until1980 int    `bson:"ate1980"`
	Until1990 int    `bson:"ate1990"`
	Until2000 int    `bson:"ate2000"`
	Until2010 int    `bson:"ate2010"`
	Total     int
}
