package db

import (
	"github.com/globalsign/mgo"
)

func WithDBSession(connStr string) (*mgo.Session, error) {
	session, err := mgo.Dial(connStr)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
