package models

import (
	_ "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Different user statuses
const (
	New = iota
	Pending
	Approved
	Suspended
	Banned
)

// User is a basic model that maps a user in the market to the Merit blockchain
type User struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	MeritAddress string        `bson:"address"`
	MeritAlias   string        `bson:"alias"`
	PublicKey    string        `bson:"publicKey"`
	Status       int           `bson:"status"` // maps to the iota above
	Verified     bool          `bson:"verified"`
	Rating       float32       `bson:"rating"`
}
