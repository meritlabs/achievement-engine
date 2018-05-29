package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meritlabs/achievement-engine/api/services"
	"github.com/meritlabs/achievement-engine/db/models"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type UsersService interface {
	CreateUserWithPassword(username, password string) (*models.User, error)
	CreateUserWithSignature(message, pubkey, signature, timestamt string, debug bool) (*models.User, error)
}

// Temp is a authentication placeholder.
// The long-term implementation should utilize public/private key encryption
// We can use the same information that exists on the Merit network, such that only we (the recipient) can decode the intended message
func Auth(service services.UsersService, l Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		debug := c.Request.Header.Get("X-Debug")
		pubkeyHex := c.Request.Header.Get("X-PubKey")
		signatureHex := c.Request.Header.Get("X-Signature")
		timestamp := c.Request.Header.Get("X-Timestamp")

		// var credentials struct {
		// 	login    string
		// 	password string
		// }
		// if err := c.ShouldBindJSON(&credentials); err != nil {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

		var user *models.User
		var err error

		fmt.Printf("url: %s", c.Request.URL.String())

		if pubkeyHex != "" && signatureHex != "" && (timestamp != "" || debug == "") {
			user, err = service.CreateUserWithSignature(c.Request.URL.String(), pubkeyHex, signatureHex, timestamp, len(debug) > 0)
			// } else if credentials.login != "" && credentials.password != "" {
			// user, err = service.CreateUserWithPassword(credentials.login, credentials.password)
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err != nil || user == nil || !user.Verified {
			fmt.Printf("Provided user is not valid! Error: %s", err.Error())
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// Set the user in the context and move onto the next function in the chain
		c.Set("user", *user)

		c.Next()
	}
}
