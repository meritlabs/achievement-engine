package stores

import "github.com/globalsign/mgo/bson"

type SessionsStore interface {
	CreateSession(userID bson.ObjectId, token string) error
	DeleteSessions(userID bson.ObjectId) error
}

func (s *Store) CreateSession(userID bson.ObjectId, token string) error {
	return nil
}

func (s *Store) DeleteSessions(userID bson.ObjectId) error {
	return nil
}
