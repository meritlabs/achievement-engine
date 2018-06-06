package dto

import (
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db/models"
)

type UserResponse struct {
	ID      bson.ObjectId `json:"id"`
	Address string        `json:"address"`
	Alias   string        `json:"alias"`
	Rating  float32       `json:"rating,omitempty"`
}

// NewUserResponseFromModel creates json response object
func NewUserResponseFromModel(m models.User) UserResponse {
	return UserResponse{
		ID:      m.ID,
		Address: m.MeritAddress,
		Alias:   m.MeritAlias,
	}
}

type TokenSessionResponse struct {
	Token string `json:"token"`
}
