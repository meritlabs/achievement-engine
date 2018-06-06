package stores

import (
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db/models"
)

type SettingsStore interface {
	GetUserSettings(userID bson.ObjectId) (*models.Settings, error)
	CreateUserSettings(userID bson.ObjectId) error
	UpdateUserSettings(userID bson.ObjectId, settings *models.Settings) error
}

func (s *Store) GetUserSettings(userID bson.ObjectId) (*models.Settings, error) {
	var settings models.Settings
	err := s.Settings.Find(bson.M{"userId": userID}).One(&settings)
	return &settings, err
}

func (s *Store) CreateUserSettings(userID bson.ObjectId) error {
	settings := models.DefaultSettings(userID)
	_, err := s.Settings.Upsert(bson.M{"userId": userID}, settings)
	return err
}

func (s *Store) UpdateUserSettings(userID bson.ObjectId, settings *models.Settings) error {
	return s.Settings.Update(
		bson.M{"userId": userID},
		bson.M{"$set": bson.M{"isSetupTrackerEnabled": settings.IsSetupTrackerEnabled}},
	)
}
