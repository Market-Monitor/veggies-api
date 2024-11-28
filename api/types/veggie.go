package types

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Veggie struct {
	M_ID bson.ObjectID `bson:"_id" json:"_id"`
	ID   string        `bson:"id" json:"id"`
	Name string        `bson:"name" json:"name"`
}

type VeggieClass struct {
	M_ID     bson.ObjectID `bson:"_id" json:"_id"`
	ID       string        `bson:"id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	ParentId string        `bson:"parentId" json:"parentId"`
}

type VeggiePrice struct {
	M_ID       bson.ObjectID `bson:"_id" json:"_id"`
	ID         string        `bson:"id" json:"id"`
	ParentId   string        `bson:"parentId" json:"parentId"`
	ParentName string        `bson:"parentName" json:"parentName"`
	Price      []int64       `bson:"price" json:"price"`
	Date       string        `bson:"date" json:"date"`
	DateISO    string        `bson:"dateISO" json:"dateISO"`
	DateUnix   int64         `bson:"dateUnix" json:"dateUnix"`
	Category   string        `bson:"category" json:"category"`
}
