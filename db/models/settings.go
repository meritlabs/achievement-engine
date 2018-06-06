package models

import (
	_ "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Settings struct {
	ID                    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserID                bson.ObjectId `bson:"userId" json:"userId"`
	IsSetupTrackerEnabled bool          `bson:"isSetupTrackerEnabled" json:"isSetupTrackerEnabled"`
}

func DefaultSettings(userId bson.ObjectId) *Settings {
	return &Settings{
		UserID:                userId,
		IsSetupTrackerEnabled: true,
	}
}
