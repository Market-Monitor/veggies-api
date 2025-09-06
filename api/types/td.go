package types

import "go.mongodb.org/mongo-driver/v2/bson"

type TradingCenter struct {
	M_ID         bson.ObjectID `bson:"_id" json:"_id"`
	Name         string        `bson:"name" json:"name"`
	LongName     string        `bson:"longName" json:"longName"`
	Slug         string        `bson:"slug" json:"slug"`
	FacebookPage string        `bson:"facebookPage" json:"facebookPage"`
}
