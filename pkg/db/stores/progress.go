package stores

import (
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/pkg/db/models/progress"
)

type ProgressStore interface {
	GetProgress(userId bson.ObjectId) (*progress.Progress, error)
	SetProgress(userId bson.ObjectId, progress progress.Progress) error
	CreateProgress(p *progress.Progress) (*progress.Progress, error)
}

func (s *Store) GetProgress(userId bson.ObjectId) (*progress.Progress, error) {
	var p progress.Progress
	err := s.Progress.Find(bson.M{"userId": userId}).One(&p)
	return &p, err
}

func (s *Store) SetProgress(userId bson.ObjectId, p progress.Progress) error {
	err := s.Progress.Update(bson.M{"userId": userId}, p)
	return err
}

func (s *Store) CreateProgress(p *progress.Progress) (*progress.Progress, error) {
	_, err := s.Progress.Upsert(bson.M{"userId": p.UserID}, p)

	if err != nil {
		return nil, err
	}

	return s.GetProgress(p.UserID)
}
