package data

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db/models"
)

func UpdateGoals(session *mgo.Session) { 
	var confirmAnInvite []models.GoalCondition
	confirmAnInvite = append(confirmAnInvite, models.GoalCondition{Name: "Confirm an invite"})
	confirmAnInvite := models.Goal{
		Slug:        1,
		Name:        "Confirm an invite",
		Description: "Confirm an invite",
		Image:       "",
		Conditions:  confirmAnInvite,
		Version:     1,
	}

	var nameYourWallet []models.GoalCondition 
	nameYourWallet = append(nameYourWallet, models.GoalCondition{Name: "Name Your Wallet"})
	nameYourWallet := models.Goal{
		Slug:        2,
		Name:        "Name Your Wallet",
		Description: "Name Your Wallet",
		Image:       "",
		Conditions:  nameYourWallet,
		Version:     1,
	}

	var hideBalance []models.GoalCondition 
	hideBalance = append(hideBalance, models.GoalCondition{Name: "Hide your wallet balance for security"})
	hideBalance := models.Goal{
		Slug:        3,
		Name:        "Hide your wallet balance for security",
		Description: "Hide your wallet balance for security",
		Image:       "",
		Conditions:  hideBalance,
		Version:     1,
	}

	var tweetShare []models.GoalCondition 
	tweetShare = append(tweetShare, models.GoalCondition{Name: "Tweet your invite link"})
	tweetShare := models.Goal{
		Slug:        4,
		Name:        "Tweet your invite link",
		Description: "Tweet your invite link",
		Image:       "",
		Conditions:  tweet,
		Version:     1,
	}

	var facebookShare []models.GoalCondition 
	facebookShare = append(facebookShare, models.GoalCondition{Name: "Share your invite link on Facebook"})
	facebookShare := models.Goal{
		Slug:        5,
		Name:        "Share your invite link on Facebook",
		Description: "Share your invite link on Facebook",
		Image:       "",
		Conditions:  facebookShare,
		Version:     1,
	}

	var enableEmailnotifications []models.GoalCondition  
	enableEmailnotifications = append(enableEmailnotifications, models.GoalCondition{Name: "Enable Email Notifications"})
	enableEmailnotifications := models.Goal{
		Slug:        6,
		Name:        "Enable Email Notifications",
		Description: "Enable Email Notifications",
		Image:       "",
		Conditions:  enableEmailnotifications,
		Version:     1,
	}

	var backupPhrase []models.GoalCondition  backupPhrase
	backupPhrase = append(backupPhrase, models.GoalCondition{Name: "Confirm your backup Phrase"})
	backupPhrase := models.Goal{
		Slug:        7,
		Name:        "Confirm your backup Phrase",
		Description: "Confirm your backup Phrase",
		Image:       "",
		Conditions:  backupPhrase,
		Version:     1,
	}

	var backupPhrase []models.GoalCondition  backupPhrase
	backupPhrase = append(backupPhrase, models.GoalCondition{Name: "Confirm your backup Phrase"})
	backupPhrase := models.Goal{
		Slug:        7,
		Name:        "Confirm your backup Phrase",
		Description: "Confirm your backup Phrase",
		Image:       "",
		Conditions:  backupPhrase,
		Version:     1,
	}

	var setPassword []models.GoalCondition  setPassword 
	setPassword = append(setPassword, models.GoalCondition{Name: "Set a Password"})
	setPassword := models.Goal{
		Slug:        8,
		Name:        "Set a Password",
		Description: "Set a Password",
		Image:       "",
		Conditions:  setPassword,
		Version:     1,
	}

	var sendMeritInvite []models.GoalCondition  sendMeritInvite   
	sendMeritInvite = append(sendMeritInvite, models.GoalCondition{Name: "Send an invite via MeritInvite"})
	sendMeritInvite := models.Goal{
		Slug:        9,
		Name:        "Send an invite via MeritInvite",
		Description: "Send an invite via MeritInvite",
		Image:       "",
		Conditions:  sendMeritInvite,
		Version:     1,
	}

	var sendMeritInvite []models.GoalCondition  sendMeritInvite   
	sendMeritInvite = append(sendMeritInvite, models.GoalCondition{Name: "Send an invite via MeritInvite"})
	sendMeritInvite := models.Goal{
		Slug:        9,
		Name:        "Send an invite via MeritInvite",
		Description: "Send an invite via MeritInvite",
		Image:       "",
		Conditions:  sendMeritInvite,
		Version:     1,
	}

	var createInviteWaitlist []models.GoalCondition  createInviteWaitlist
	createInviteWaitlist = append(createInviteWaitlist, models.GoalCondition{Name: "Add friends to your invite waitlist"})
	createInviteWaitlist := models.Goal{
		Slug:        10,
		Name:        "Add friends to your invite waitlist",
		Description: "Add friends to your invite waitlist",
		Image:       "",
		Conditions:  createInviteWaitlist,
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
