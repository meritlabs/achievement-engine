package models

import (
	_ "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type GoalCondition struct {
	Slug        int    `bson:"slug" json:"slug"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	GoalSlug    int    `bson:"goalSlug" json:"goalSlug"`
}

type Goal struct {
	ID          bson.ObjectId   `bson:"_id,omitempty" json:"id"`
	Slug        int             `bson:"slug" json:"slug"`
	Name        string          `bson:"name" json:"name"`
	Description string          `bson:"description" json:"description"`
	Image       string          `bson:"image" json:"image"`
	Conditions  []GoalCondition `bson:"conditions" json:"conditions"`
	Version     int             `bson:"version" json:"version"`
}
