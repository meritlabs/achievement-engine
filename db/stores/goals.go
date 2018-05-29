package stores

import (
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db/models"
)

type GoalsStore interface {
	ListGoals() ([]models.Goal, error)
}

func (s *Store) ListGoals() ([]models.Goal, error) {
	var goals []models.Goal
	err := s.Goals.Find(bson.M{}).All(&goals)

	if err != nil {
		return nil, err
	}

	return goals, nil
}
