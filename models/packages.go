package models

import (
	"time"
)

type Location struct {
	Address     string    `json:"address" bson:"address"`
	Zipcode     int64     `json:"zipcode" bson:"zipcode"`
	Extnum      int       `json:"ext_num" bson:"ext_num"`
	IntNum      int       `json:"int_num" bson:"int_num"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type Packages struct {
	Identifier string    `json:"identifier" bson:"identifier"`
	Size       string    `json:"size" bson:"size"`
	Weight     float64   `json:"weight" bson:"weight"`
	Amount     int       `json:"amount" bson:"amount"`
	Status     string    `json:"status" bson:"status"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	Origin     *Location `json:"origin" bson:"origin"`
	Destiny    *Location `json:"destiny" bson:"destiny"`
}
