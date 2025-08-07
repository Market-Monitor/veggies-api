package types

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type AllVeggieWithClasses struct {
	TradingCenters []TradingCenter `json:"tradingCenters"`
	Veggies        []AllVeggies    `json:"veggies"`
}

type AllVeggies struct {
	Data Veggie `json:"data"`

	Classes []VeggieClass `json:"classes"`
}

type FilterVeggies struct {
	ImageSource      string                    `bson:"imageSource" json:"imageSource"`
	ParentId         string                    `bson:"parentId" json:"parentId"`
	ParentName       string                    `bson:"parentName" json:"parentName"`
	Category         string                    `bson:"category" json:"category"`
	TradingCenter    string                    `bson:"tradingCenter" json:"tradingCenter"`
	PriceUnit        string                    `bson:"priceUnit" json:"priceUnit"`
	LatestUpdateDate int64                     `bson:"latestUpdateDate" json:"latestUpdateDate"`
	Classes          []LatestHistoryPriceClass `bson:"classes" json:"classes"`
}

type Veggie struct {
	M_ID          bson.ObjectID `bson:"_id" json:"_id"`
	ID            string        `bson:"id" json:"id"`
	Name          string        `bson:"name" json:"name"`
	TradingCenter string        `bson:"tradingCenter" json:"tradingCenter"`
	ImageURL      string        `bson:"imageUrl" json:"imageUrl"`
	ImageSource   string        `bson:"imageSource" json:"imageSource"`
	PriceUnit     string        `bson:"priceUnit" json:"priceUnit"`
}

type VeggieClass struct {
	M_ID          bson.ObjectID `bson:"_id" json:"_id"`
	ID            string        `bson:"id" json:"id"`
	Name          string        `bson:"name" json:"name"`
	ParentId      string        `bson:"parentId" json:"parentId"`
	TradingCenter string        `bson:"tradingCenter" json:"tradingCenter"`
}

type VeggiePrice struct {
	M_ID          bson.ObjectID `bson:"_id" json:"_id"`
	ID            string        `bson:"id" json:"id"`
	Name          string        `bson:"name" json:"name"`
	ParentId      string        `bson:"parentId" json:"parentId"`
	ParentName    string        `bson:"parentName" json:"parentName"`
	Price         []int64       `bson:"price" json:"price"`
	Date          string        `bson:"date" json:"date"`
	DateISO       string        `bson:"dateISO" json:"dateISO"`
	DateUnix      int64         `bson:"dateUnix" json:"dateUnix"`
	Category      string        `bson:"category" json:"category"`
	TradingCenter string        `bson:"tradingCenter" json:"tradingCenter"`
}

type LatestHistoryPrice struct {
	ImageSource   string                    `bson:"imageSource" json:"imageSource"`
	ParentId      string                    `bson:"parentId" json:"parentId"`
	ParentName    string                    `bson:"parentName" json:"parentName"`
	Category      string                    `bson:"category" json:"category"`
	TradingCenter string                    `bson:"tradingCenter" json:"tradingCenter"`
	Classes       []LatestHistoryPriceClass `bson:"classes" json:"classes"`
	PriceUnit     string                    `bson:"priceUnit" json:"priceUnit"`
}

type LatestHistoryPriceClass struct {
	ID    string  `bson:"id" json:"id"`
	Name  string  `bson:"name" json:"name"`
	Price []int64 `bson:"price" json:"price"`
}
