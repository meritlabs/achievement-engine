package data

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db/models"
)

func UpdateGoals(session *mgo.Session) {
	var confirmAnInviteConditions []models.GoalCondition
	confirmAnInviteConditions = append(confirmAnInviteConditions, models.GoalCondition{Name: "Confirm an invite"})
	confirmAnInvite := models.Goal{
		Slug:        1,
		Name:        "Confirm an invite",
		Description: "Confirm an invite",
		Image:       "",
		Conditions:  confirmAnInviteConditions,
		Version:     1,
	}

	var nameYourWalletConditions []models.GoalCondition
	nameYourWalletConditions = append(nameYourWalletConditions, models.GoalCondition{Name: "Name Your Wallet"})
	nameYourWallet := models.Goal{
		Slug:        2,
		Name:        "Name Your Wallet",
		Description: "Name Your Wallet",
		Image:       "",
		Conditions:  nameYourWalletConditions,
		Version:     1,
	}

	var hideBalanceConditions []models.GoalCondition
	hideBalanceConditions = append(hideBalanceConditions, models.GoalCondition{Name: "Hide your wallet balance for security"})
	hideBalance := models.Goal{
		Slug:        3,
		Name:        "Hide your wallet balance for security",
		Description: "Hide your wallet balance for security",
		Image:       "",
		Conditions:  hideBalanceConditions,
		Version:     1,
	}

	var tweetShareConditions []models.GoalCondition
	tweetShareConditions = append(tweetShareConditions, models.GoalCondition{Name: "Tweet your invite link"})
	tweetShare := models.Goal{
		Slug:        4,
		Name:        "Tweet your invite link",
		Description: "Tweet your invite link",
		Image:       "",
		Conditions:  tweetShareConditions,
		Version:     1,
	}

	var facebookShareConditions []models.GoalCondition
	facebookShareConditions = append(facebookShareConditions, models.GoalCondition{Name: "Share your invite link on Facebook"})
	facebookShare := models.Goal{
		Slug:        5,
		Name:        "Share your invite link on Facebook",
		Description: "Share your invite link on Facebook",
		Image:       "",
		Conditions:  facebookShareConditions,
		Version:     1,
	}

	var enableEmailnotificationsConditions []models.GoalCondition
	enableEmailnotificationsConditions = append(enableEmailnotificationsConditions, models.GoalCondition{Name: "Enable Email Notifications"})
	enableEmailnotifications := models.Goal{
		Slug:        6,
		Name:        "Enable Email Notifications",
		Description: "Enable Email Notifications",
		Image:       "",
		Conditions:  enableEmailnotificationsConditions,
		Version:     1,
	}

	var backupPhraseConditions []models.GoalCondition
	backupPhraseConditions = append(backupPhraseConditions, models.GoalCondition{Name: "Confirm your backup Phrase"})
	backupPhrase := models.Goal{
		Slug:        7,
		Name:        "Confirm your backup Phrase",
		Description: "Confirm your backup Phrase",
		Image:       "",
		Conditions:  backupPhraseConditions,
		Version:     1,
	}

	var setPasswordConditions []models.GoalCondition
	setPasswordConditions = append(setPasswordConditions, models.GoalCondition{Name: "Set a Password"})
	setPassword := models.Goal{
		Slug:        8,
		Name:        "Set a Password",
		Description: "Set a Password",
		Image:       "",
		Conditions:  setPasswordConditions,
		Version:     1,
	}

	var sendMeritInviteConditions []models.GoalCondition
	sendMeritInviteConditions = append(sendMeritInviteConditions, models.GoalCondition{Name: "Send an invite via MeritInvite"})
	sendMeritInvite := models.Goal{
		Slug:        9,
		Name:        "Send an invite via MeritInvite",
		Description: "Send an invite via MeritInvite",
		Image:       "",
		Conditions:  sendMeritInviteConditions,
		Version:     1,
	}

	var createInviteWaitlistConditions []models.GoalCondition
	createInviteWaitlistConditions = append(createInviteWaitlistConditions, models.GoalCondition{Name: "Add friends to your invite waitlist"})
	createInviteWaitlist := models.Goal{
		Slug:        10,
		Name:        "Add friends to your invite waitlist",
		Description: "Add friends to your invite waitlist",
		Image:       "",
		Conditions:  createInviteWaitlistConditions,
		Version:     1,
	}

	db := session.DB("achievement-engine").C("goals")
	db.Upsert(bson.M{"slug": confirmAnInvite.Slug}, confirmAnInvite)
	db.Upsert(bson.M{"slug": nameYourWallet.Slug}, nameYourWallet)
	db.Upsert(bson.M{"slug": hideBalance.Slug}, hideBalance)
	db.Upsert(bson.M{"slug": tweetShare.Slug}, tweetShare)
	db.Upsert(bson.M{"slug": facebookShare.Slug}, facebookShare)
	db.Upsert(bson.M{"slug": enableEmailnotifications.Slug}, enableEmailnotifications)
	db.Upsert(bson.M{"slug": backupPhrase.Slug}, backupPhrase)
	db.Upsert(bson.M{"slug": setPassword.Slug}, setPassword)
	db.Upsert(bson.M{"slug": sendMeritInvite.Slug}, sendMeritInvite)
	db.Upsert(bson.M{"slug": createInviteWaitlist.Slug}, createInviteWaitlist)

}
