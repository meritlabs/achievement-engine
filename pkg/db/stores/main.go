package stores

import (
	"github.com/globalsign/mgo"

	"github.com/meritlabs/achievement-engine/pkg/db"
)

type Store struct {
	session      *mgo.Session
	Users        *mgo.Collection
	Sessions     *mgo.Collection
	Settings     *mgo.Collection
	Progress     *mgo.Collection
}

func InitStore(connStr string) *Store {
	session, err := db.WithDBSession(connStr)

	if err != nil {
		panic(err)
	}

	db := session.DB("achievement-engine")
	users := db.C("users")
	sessions := db.C("sessions")
	settings := db.C("settings")
	progress := db.C("progress")

	return &Store{
		session:      session,
		Users:        users,
		Sessions:     sessions,
		Settings:     settings,
		Progress:     progress,
	}
}

func (store *Store) ShutDownStore() {
	store.session.Close()
}
