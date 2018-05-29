package data

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db/models"
)

func UpdateGoals(session *mgo.Session) {
	var whoAmIConditions []models.GoalCondition
	whoAmIConditions = append(whoAmIConditions, models.GoalCondition{Name: "Find alias and paste into answer field"})
	whoAmI := models.Goal{
		Slug:        1,
		Name:        "Who am I",
		Description: "Teach users to find thei aliases and explaixn what is it",
		Image:       "",
		Conditions:  whoAmIConditions,
		Version:     1,
	}

	var backupRookieConditions []models.GoalCondition
	backupRookieConditions = append(backupRookieConditions, models.GoalCondition{Name: "Find mnemonic phrase and paste it into answer field"})
	backupRookie := models.Goal{
		Slug:        2,
		Name:        "BackUp rookie",
		Description: "Teach users to backup their wallets",
		Image:       "",
		Conditions:  backupRookieConditions,
		Version:     1,
	}

	var beginnerBrokerConditions []models.GoalCondition
	beginnerBrokerConditions = append(beginnerBrokerConditions,
		models.GoalCondition{
			Name: "Login into market",
		},
		models.GoalCondition{
			Name: "Create one or more BID/ASK",
		},
	)
	beginnerBroker := models.Goal{
		Slug:        3,
		Name:        "Beginner brocker",
		Description: "Teach users how to Buy/Sell MRT via Market",
		Image:       "",
		Conditions:  beginnerBrokerConditions,
		Version:     1,
	}

	db := session.DB("achievement-engine").C("goals")
	db.Upsert(bson.M{"slug": whoAmI.Slug}, whoAmI)
	db.Upsert(bson.M{"slug": backupRookie.Slug}, backupRookie)
	db.Upsert(bson.M{"slug": beginnerBroker.Slug}, beginnerBroker)
}
