package db

import (
	"github.com/globalsign/mgo"
)

func WithDBSession() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
