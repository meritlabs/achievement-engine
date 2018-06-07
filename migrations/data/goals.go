package data

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db/models"
)

func UpdateGoals(session *mgo.Session) {

	var createFirstWallet []models.GoalCondition
	createFirstWallet = append(createFirstWallet, models.GoalCondition{Name: "Create Wallet"})
	createFirstWallet := models.Goal{
		Slug:        0,
		Route: 			 "wallets",
		Name:        "Create Wallet",
		Description: "Create your first wallet",
		Title				 "Create",
		LinkTitle:	 "Wallet",
		Image:       "",
		Conditions:  createFirstWallet,
		Version:     1,
	}

	var nameYourWalletConditions []models.GoalCondition
	nameYourWalletConditions = append(nameYourWalletConditions, models.GoalCondition{Name: "Name Your Wallet"})
	nameYourWallet := models.Goal{
		Slug:        1,
		Route: 			 "wallets",
		Name:        "Name Your Wallet",
		Description: "For better experiance you can rename your wallet",
		Title				 "Name Your",
		LinkTitle:	 "Wallet",
		Image:       "",
		Conditions:  nameYourWalletConditions,
		Version:     1,
	}

	var hideBalanceConditions []models.GoalCondition
	hideBalanceConditions = append(hideBalanceConditions, models.GoalCondition{Name: "Hide your wallet balance"})
	hideBalance := models.Goal{
		Slug:        2,
		Route: 			 "wallets",
		Name:        "Hide your wallet balance",
		Description: "Hide your wallet balance for security",
		Title				 "Hide your",
		LinkTitle:	 "wallet balance",
		Image:       "",
		Conditions:  hideBalanceConditions,
		Version:     1,
	}

	var confirmAnInviteConditions []models.GoalCondition
	confirmAnInviteConditions = append(confirmAnInviteConditions, models.GoalCondition{Name: "Share your invite code"})
	confirmAnInvite := models.Goal{
		Slug:        3,
		Route: 			 "invite-action",
		Name:        "Share your invite code",
		Description: "Spread to the World your invite code and exted your community...",
		Title				 "Share your",
		LinkTitle:	 "invite code",
		Image:       "",
		Conditions:  confirmAnInviteConditions,
		Version:     1,
	}

	var createInviteWaitlistConditions []models.GoalCondition
	createInviteWaitlistConditions = append(createInviteWaitlistConditions, models.GoalCondition{Name: "Add friends to your invite waitlist"})
	createInviteWaitlist := models.Goal{
		Slug:        4,
		Route: 			 "invites/requests",
		Name:        "Add friends to your invite waitlist",
		Description: "Add friends to your invite waitlist",
		Title				 "Add friends to your",
		LinkTitle:	 "invite waitlist",
		Image:       "",
		Conditions:  createInviteWaitlistConditions,
		Version:     1,
	}

	var backupPhraseConditions []models.GoalCondition
	backupPhraseConditions = append(backupPhraseConditions, models.GoalCondition{Name: "Confirm your backup Phrase"})
	backupPhrase := models.Goal{
		Slug:        5,
		Route: 			 "wallets",
		Name:        "Confirm your backup Phrase",
		Description: "Wallet Backup Phrase it's easiest way to ",
		Image:       "",
		Conditions:  backupPhraseConditions,
		Version:     1,
	}

	var setPasswordConditions []models.GoalCondition
	setPasswordConditions = append(setPasswordConditions, models.GoalCondition{Name: "Set a advanced Wallet Password"})
	setPassword := models.Goal{
		Slug:        6,
		Route: 			 "wallets",
		Name:        "Set a advanced Wallet Password",
		Description: "We are very care about your security, you can setup advanced password for protect your Wallet Backup Phrase",
		Title				 "Set a advanced",
		LinkTitle:	 "Wallet Password",
		Image:       "",
		Conditions:  setPasswordConditions,
		Version:     1,
	}

	var enableEmailnotificationsConditions []models.GoalCondition
	enableEmailnotificationsConditions = append(enableEmailnotificationsConditions, models.GoalCondition{Name: "Enable Email Notifications"})
	enableEmailnotifications := models.Goal{
		Slug:        7,
		Route: 			 "settings",
		Name:        "Enable Email Notifications",
		Description: "Enable Email Notifications",
		Title				 "Enable",
		LinkTitle:	 "Notifications",
		Image:       "",
		Conditions:  enableEmailnotificationsConditions,
		Version:     1,
	}

	

	db := session.DB("achievement-engine").C("goals")
	db.Upsert(bson.M{"slug": createFirstWallet.Slug}, createFirstWallet)
	db.Upsert(bson.M{"slug": confirmAnInvite.Slug}, confirmAnInvite)
	db.Upsert(bson.M{"slug": nameYourWallet.Slug}, nameYourWallet)
	db.Upsert(bson.M{"slug": hideBalance.Slug}, hideBalance)
	db.Upsert(bson.M{"slug": enableEmailnotifications.Slug}, enableEmailnotifications)
	db.Upsert(bson.M{"slug": backupPhrase.Slug}, backupPhrase)
	db.Upsert(bson.M{"slug": setPassword.Slug}, setPassword)
	db.Upsert(bson.M{"slug": sendMeritInvite.Slug}, sendMeritInvite)
	db.Upsert(bson.M{"slug": createInviteWaitlist.Slug}, createInviteWaitlist)

}
