package progress

import (
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/pkg/db/models/goal"
)

const (
	Incomplete = iota
	Inprogress
	Complete
)

type TaskProgress struct {
	Slug   goal.TaskSlug `bson:"slug" json:"slug"`
	Status int           `bson:"status" json:"status"`
}

type Progress struct {
	ID     bson.ObjectId  `bson:"_id,omitempty" json:"id"`
	UserID bson.ObjectId  `bson:"userId" json:"userId"`
	Tasks  []TaskProgress `bson:"tasks" json:"tasks"`
}
