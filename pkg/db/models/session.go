package models

import (
	_ "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Session struct {
	ID     bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserID bson.ObjectId `bson:"userId" json:"userId"`
	Token  string        `bson:"token" json:"token"`
}
