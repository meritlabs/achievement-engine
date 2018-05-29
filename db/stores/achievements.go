package stores

import (
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db/models"
)

type AchievementsStore interface {
	GetAchievementsForUser(userID bson.ObjectId) (*[]models.Achievement, error)
	GetAchievementForUser(userID bson.ObjectId, achivementId bson.ObjectId) (*models.Achievement, error)
	CopyAchievementsFromGoals(userID bson.ObjectId, goals []models.Goal) (*[]models.Achievement, error)
}

func (s *Store) GetAchievementsForUser(userID bson.ObjectId) (*[]models.Achievement, error) {
	var achievements []models.Achievement
	err := s.Achievements.Find(bson.M{"userId": userID}).All(&achievements)
	return &achievements, err
}

func (s *Store) GetAchievementForUser(userID bson.ObjectId, achivementId bson.ObjectId) (*models.Achievement, error) {
	var achievement models.Achievement
	err := s.Achievements.Find(bson.M{"userId": userID, "_id": achivementId}).One(&achievement)
	return &achievement, err
}

func (s *Store) CopyAchievementsFromGoals(userID bson.ObjectId, goals []models.Goal) (*[]models.Achievement, error) {
	var achievements []models.Achievement
	for _, goal := range goals {
		a := models.Achievement{}
		a.FromGoal(&goal)
		a.UserID = userID

		err := s.Achievements.Insert(a)

		if err != nil {
			return nil, err
		}

		achievements = append(achievements, a)
	}

	return &achievements, nil
}
