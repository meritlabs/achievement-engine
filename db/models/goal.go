package models

import (
	_ "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type GoalCondition struct {
	Slug int    `bson:"slug" json:"slug"`
	Name string `bson:"name" json:"name"`
}

type Goal struct {
	ID          bson.ObjectId   `bson:"_id,omitempty" json:"id"`
	Slug        int             `bson:"slug" json:"slug"`
	Route       string          `bson:"route" json:"route"`
	Name        string          `bson:"name" json:"name"`
	Description string          `bson:"description" json:"description"`
	Title       string          `bson:"title" json:"title"`
	LinkTitle   string          `bson:"linkTitle" json:"linkTitle"`
	Image       string          `bson:"image" json:"image"`
	Conditions  []GoalCondition `bson:"conditions" json:"conditions"`
	Version     int             `bson:"version" json:"version"`
}
