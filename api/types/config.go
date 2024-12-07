package types

type Configuration struct {
	LatestDataDate string `bson:"latestDataDate" json:"latestDataDate"`
	ConfigId       int64  `bson:"configId" json:"configId"`
}
