package stores

import "github.com/meritlabs/achievement-engine/db/models"

type GoalsStore interface {
	ListGoals() ([]models.Goal, error)
}

func (s *Store) ListGoals() ([]models.Goal, error) {
	return nil, nil
}
