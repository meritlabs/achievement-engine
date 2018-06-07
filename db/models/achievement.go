package models

import (
	_ "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	NotStarted = iota
	InProgress
	Done
)

type AchievementCondition struct {
	Slug   int    `bson:"slug" json:"slug"`
	Name   string `bson:"name" json:"name"`
	Status int    `bson:"status" json:"status"`
}

func (ac *AchievementCondition) FromGoalCondition(gc *GoalCondition) {
	ac.Name = gc.Name
	ac.Slug = gc.Slug
	ac.Status = NotStarted
}

type Achievement struct {
	ID          bson.ObjectId          `bson:"_id,omitempty" json:"id"`
	GoalID      bson.ObjectId          `bson:"goalId" json:"goalId"`
	UserID      bson.ObjectId          `bson:"userId" json:"userId"`
	Slug        int                    `bson:"slug" json:"slug"`
	Route       string                 `bson:"route" json:"route"`
	Name        string                 `bson:"name" json:"name"`
	Description string                 `bson:"description" json:"description"`
	Title       string                 `bson:"title" json:"title"`
	LinkTitle   string                 `bson:"linkTitle" json:"linkTitle"`
	Image       string                 `bson:"image" json:"image"`
	Conditions  []AchievementCondition `bson:"conditions" json:"conditions"`
	Version     int                    `bson:"version" json:"version"`
	Status      int                    `bson:"status" json:"status"`
}

func (a *Achievement) FromGoal(goal *Goal) {
	a.Slug = goal.Slug
	a.Route = goal.Route
	a.Name = goal.Name
	a.Description = goal.Description
	a.Title = goal.Title
	a.LinkTitle = goal.LinkTitle
	a.GoalID = goal.ID
	a.Image = goal.Image
	a.Version = goal.Version
	a.Status = NotStarted

	var conditions []AchievementCondition
	for _, c := range goal.Conditions {
		ac := AchievementCondition{}
		ac.FromGoalCondition(&c)
		conditions = append(conditions, ac)
		ac.Status = NotStarted
	}
	a.Conditions = conditions
}
