package stores

import (
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db/models"
)

type UsersStore interface {
	GetUser(id bson.ObjectId) (*models.User, error)
	ListUsers() (*[]models.User, error)
	UserByPublicKey(pubkey string) (*models.User, error)
	CreateUser(user *models.User) error
	CreateUserByAddress(address string, user *models.User) error
}

// User returns a user by id
func (s *Store) GetUser(id bson.ObjectId) (*models.User, error) {
	var user models.User
	err := s.Users.Find(bson.M{"_id": id}).One(&user)
	return &user, err
}

// UserByPublicKey returns a user by public key hex string
func (s *Store) UserByPublicKey(pubkey string) (*models.User, error) {
	var user models.User
	err := s.Users.Find(bson.M{"publicKey": pubkey}).One(&user)
	return &user, err
}

// Users returns a list of users
func (s *Store) ListUsers() (*[]models.User, error) {
	var users []models.User
	err := s.Users.Find(bson.M{}).All(&users)
	return &users, err
}

// CreateUser store a user in a db
func (s *Store) CreateUser(user *models.User) error {
	return s.Users.Insert(user)
}

// CreateUserByAddress creates a user by address if not exists
func (s *Store) CreateUserByAddress(address string, user *models.User) error {
	_, err := s.Users.Upsert(bson.M{"address": address}, user)
	return err
}
