package stores

import (
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db/models"
)

type SessionsStore interface {
	GetSession(token string) (*models.Session, error)
	CreateSession(userID bson.ObjectId, token string) error
	DeleteSessions(userID bson.ObjectId) error
}

func (s *Store) GetSession(token string) (*models.Session, error) {
	var session models.Session
	err := s.Sessions.Find(bson.M{"token": token}).One(&session)
	return &session, err
}

func (s *Store) CreateSession(userID bson.ObjectId, token string) error {
	session := models.Session{
		UserID: userID,
		Token:  token,
	}
	return s.Sessions.Insert(&session)
}

func (s *Store) DeleteSessions(userID bson.ObjectId) error {
	_, err := s.Sessions.RemoveAll(bson.M{"userId": userID})
	return err
}
