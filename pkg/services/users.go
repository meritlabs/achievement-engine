package services

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/meritlabs/achievement-engine/pkg/db/models"
	"github.com/meritlabs/achievement-engine/pkg/db/stores"
	"github.com/meritlabs/achievement-engine/pkg/utils"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

// used as a prefix in signmessage method in meritd
const msgMagic = "Merit Signed Message:\n"

// UsersService
type UsersService struct {
	NetParams     chaincfg.Params
	BCClient      Client
	UsersStore    stores.UsersStore
	SessionsStore stores.SessionsStore
	SettingsStore stores.SettingsStore
	ProgressStore stores.ProgressStore
}

/// NEEW
func ParsePubKey(pubkeyHex string) (*btcec.PublicKey, error) {
	// Decode hex-encoded pubkey.
	pubkeyBytes, err := hex.DecodeString(pubkeyHex)
	if err != nil {
		return nil, err
	}

	return btcec.ParsePubKey(pubkeyBytes, btcec.S256())
}

// CheckSignature checks that provided message was signed with the given key
func CheckSignature(message, pubkeyHex, signatureHex, timestamp string, debug bool) error {
	msg := message

	// append timestamp to the message if not in debug mode.
	// useful for testing as timestamp changes signature every second
	if !debug {
		msg += timestamp
	}

	var buf bytes.Buffer
	wire.WriteVarString(&buf, 0, msgMagic)
	wire.WriteVarString(&buf, 0, msg)

	messageHash := chainhash.DoubleHashB(buf.Bytes())

	pubkey, err := ParsePubKey(pubkeyHex)
	if err != nil {
		return err
	}

	// Decode hex-encoded signature.
	sigBytes, err := base64.StdEncoding.DecodeString(signatureHex)
	if err != nil {
		return err
	}

	pk, _, err := btcec.RecoverCompact(btcec.S256(), sigBytes, messageHash)

	if err != nil {
		return err
	}

	if !pubkey.IsEqual(pk) {
		return errors.New("invalid signature")
	}

	return nil
}

// CreateUserWithSignature creates a user with provided pubkey and signature
// TODO: replace pubkey with address
func (s *UsersService) CreateUserWithSignature(message, pubkeyHex, signatureHex, timestamp string, debug bool) (*models.User, error) {
	if err := CheckSignature(message, pubkeyHex, signatureHex, timestamp, debug); err != nil {
		return nil, err
	}

	pubkey, err := ParsePubKey(pubkeyHex)
	if err != nil {
		return nil, err
	}

	pubkeyHash := btcutil.Hash160(pubkey.SerializeCompressed())

	address, err := btcutil.NewAddressPubKeyHash(pubkeyHash, &s.NetParams)
	addressStr := address.String()

	// Let's look up the user, and insert their information into the context as it flows through the request/response chain.
	user := models.User{MeritAddress: addressStr}
	if err := s.UsersStore.CreateUserByAddress(addressStr, &user); err != nil {
		return nil, err
	}

	if user.Status != models.Approved && user.Status != models.Banned || user.Status == models.Pending {
		// do it synchronous as we need to know that address is valid and confirmed
		if err := s.getUserFromBlockchain(&user, pubkeyHex); err != nil {
			return nil, err
		}
	}

	s.setUpUserData(&user)

	return &user, nil
}

// CreateUserWithPassword create a user with a provided ligin/password pair
func (s *UsersService) CreateUserWithPassword(username, password string) (*models.User, error) {
	fmt.Printf("Signup with login/password\n")

	return nil, nil
}

// CreateSession create new session for a give user
func (s *UsersService) CreateSession(user models.User) (string, error) {
	s.SessionsStore.DeleteSessions(user.ID)

	token, err := utils.GenerateRandomString(32)
	if err != nil {
		return "", err
	}

	if err := s.SessionsStore.CreateSession(user.ID, token); err != nil {
		return "", err
	}

	return token, nil
}

func (s *UsersService) getUserFromBlockchain(user *models.User, pubkey string) error {
	if len(pubkey) == 0 {
		return errors.New("empty pubkey")
	}

	addressInfo, err := s.BCClient.ValidateAddress(user.MeritAddress)

	if len(addressInfo.Address) == 0 {
		return errors.New("user not found in blockchain")
	}

	if err != nil {
		return err
	}

	if addressInfo.IsValid {
		user.Verified = true
	}

	if addressInfo.Beaconed > 0 && addressInfo.Confirmed == 0 {
		user.Status = models.Pending
	}

	if addressInfo.Beaconed > 0 && addressInfo.Confirmed > 0 {
		user.Status = models.Approved
	}

	user.MeritAlias = addressInfo.Alias
	user.PublicKey = pubkey

	return s.UsersStore.CreateUser(user)
}

func (s *UsersService) setUpUserData(user *models.User) error {
	settings, err := s.SettingsStore.GetUserSettings(user.ID)
	if err != nil {
		return err
	}

	if settings == nil {
		s.SettingsStore.CreateUserSettings(user.ID)
	}

	//achievements, err := s.AchievementsStore.GetAchievementsForUser(user.ID)
	//if err != nil {
	//	return err
	//}
	//
	//if len(*achievements) > 0 {
	//	return nil
	//}
	//
	//goals, err := s.GoalsStore.ListGoals()
	//if err != nil {
	//	return err
	//}
	//
	//_, err = s.AchievementsStore.CopyAchievementsFromGoals(user.ID, goals)
	//if err != nil {
	//	return err
	//}

	return nil
}
