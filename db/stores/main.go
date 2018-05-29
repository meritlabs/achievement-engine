package stores

import (
	"github.com/globalsign/mgo"

	"github.com/meritlabs/achievement-engine/db"
)

type Store struct {
	session      *mgo.Session
	Achievements *mgo.Collection
	Goals        *mgo.Collection
	Users        *mgo.Collection
	Sessions     *mgo.Collection
}

func InitStore() *Store {
	session, err := db.WithDBSession()

	if err != nil {
		panic(err)
	}

	db := session.DB("achievement-engine")
	achievements := db.C("achievements")
	goals := db.C("goals")
	users := db.C("users")
	sessions := db.C("sessions")

	return &Store{
		session:      session,
		Achievements: achievements,
		Goals:        goals,
		Users:        users,
		Sessions:     sessions,
	}
}

func (store *Store) ShutDownStore() {
	store.session.Close()
}
