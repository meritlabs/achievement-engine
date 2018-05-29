package main

import (
	"github.com/meritlabs/achievement-engine/db"
	"github.com/meritlabs/achievement-engine/migrations/data"
)

func main() {
	session, err := db.WithDBSession()

	if err != nil {
		panic(err)
	}
	defer session.Close()

	data.UpdateGoals(session)
}
