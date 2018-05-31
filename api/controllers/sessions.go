package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meritlabs/achievement-engine/db/models"
	dto "github.com/meritlabs/achievement-engine/api/models"
	"github.com/meritlabs/achievement-engine/api/services"
)

func CreateSession(s services.UsersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		debug := c.Request.Header.Get("X-Debug")
		pubkeyHex := c.Request.Header.Get("X-PubKey")
		signatureHex := c.Request.Header.Get("X-Signature")
		timestamp := c.Request.Header.Get("X-Timestamp")

		var user *models.User
		var err error

		if pubkeyHex == "" || signatureHex == "" || (timestamp == "" && debug == "") {
			c.Error(dto.BadRequestError{Message: "authorization parameters missing"})
			return
		}

		user, err = s.CreateUserWithSignature(c.Request.URL.String(), pubkeyHex, signatureHex, timestamp, len(debug) > 0)

		if err != nil || user == nil || user.Status != models.Approved {
			fmt.Printf("Access denied: err: %v \n user: %v \n", err, user)
			c.Error(dto.ForbiddenError{})
			return
		}

		token, err := s.CreateSession(*user)

		if err != nil {
			fmt.Printf("Access denied: err: %v \n token: %v \n", err, token)
			c.Error(dto.ForbiddenError{})
			return
		}

		c.JSON(http.StatusOK, dto.NewSessionResponseFromModel(token, *user))
	}
}
