package data

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db/models"
)

func UpdateGoals(session *mgo.Session) {

	var creatorConditions []models.GoalCondition
	creatorConditions = append(creatorConditions, models.GoalCondition{
		Name:        "Create wallet",
		Description: "",
		GoalSlug:    1})
	creatorConditions = append(creatorConditions, models.GoalCondition{
		Name:        "Unlock wallet",
		Description: "",
		GoalSlug:    2})
	creator := models.Goal{
		Slug:        1,
		Name:        "Creator",
		Description: "Create and unlock your first wallet.",
		Image:       "achi-creator",
		Conditions:  creatorConditions,
		Version:     1,
	}

	var fastStarterConditions []models.GoalCondition
	fastStarterConditions = append(fastStarterConditions, models.GoalCondition{
		Name:        "Invite Your Friends to Merit!",
		Description: "Share your alias with your friends!  Your alias can be used as their invite code when they create a wallet.",
		GoalSlug:    1})
	fastStarterConditions = append(fastStarterConditions, models.GoalCondition{
		Name:        "Add a friend to your invite waitlist",
		Description: "You can populate your invite waitlist at any time, even when you don’t have available invites! Just share your alias with your friends and have them use it as their invite code when they create their Merit Wallet.",
		GoalSlug:    2})
	fastStarterConditions = append(fastStarterConditions, models.GoalCondition{
		Name:        "Mine an invite",
		Description: "Invites are randomly distributed among the Merit community with every block that is mined.  Keep an eye out for new invites in your wallet!",
		GoalSlug:    3})
	fastStarterConditions = append(fastStarterConditions, models.GoalCondition{
		Name:        "Confirm an invite request",
		Description: "Confirm the invite requests pending in your invite waitlist! Remember, invites are scarce!",
		GoalSlug:    4})
	fastStarter := models.Goal{
		Slug:        2,
		Name:        "Fast Starter",
		Description: "Create and unlock your first wallet.",
		Image:       "achi-start",
		Conditions:  fastStarterConditions,
		Version:     1,
	}

	var tycoonConditions []models.GoalCondition
	tycoonConditions = append(tycoonConditions, models.GoalCondition{
		Name:        "Send a friend Merit using MeritMoney",
		Description: "MeritMoney lets you add a password to the transaction or cancel it at any time up until the transaction is accepted. ",
		GoalSlug:    1})
	tycoonConditions = append(tycoonConditions, models.GoalCondition{
		Name:        "Chat with the Merit team on Discord or Telegram",
		Description: "Join us on Discord or Telegram to hear about the latest updates, news, and get answers to any questions you might have about Merit.",
		GoalSlug:    2})
	tycoonConditions = append(tycoonConditions, models.GoalCondition{
		Name:        "Receive merit from friend or community",
		Description: "Your balance should be more than 1MRT.",
		GoalSlug:    3})
	tycoonConditions = append(tycoonConditions, models.GoalCondition{
		Name:        "Get SMS notifications",
		Description: "We can text you and let you know when you get new invites and when you receive MRT.  Just add your number.  ",
		GoalSlug:    4})
	tycoon := models.Goal{
		Slug:        3,
		Name:        "Merit Tycoon",
		Description: "Do more activity in the Merit World",
		Image:       "achi-tycoon",
		Conditions:  tycoonConditions,
		Version:     1,
	}

	var growthMasterConditions []models.GoalCondition
	growthMasterConditions = append(growthMasterConditions, models.GoalCondition{
		Name:        "Earn your first Growth Reward",
		Description: "Merit believes in rewarding community members for growing the network.  Adding active users who mine, sell, and invite others will help you earn more Growth Rewards.",
		GoalSlug:    1})
	growthMaster := models.Goal{
		Slug:        4,
		Name:        "Growth Master",
		Description: "Growth your community!",
		Image:       "achi-growth-master",
		Conditions:  growthMasterConditions,
		Version:     1,
	}

	db := session.DB("achievement-engine").C("goals")
	db.Upsert(bson.M{"slug": creator.Slug}, creator)
	db.Upsert(bson.M{"slug": fastStarter.Slug}, creator)
	db.Upsert(bson.M{"slug": tycoon.Slug}, creator)
	db.Upsert(bson.M{"slug": growthMaster.Slug}, creator)

}
