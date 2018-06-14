package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meritlabs/achievement-engine/pkg/db/models"
	"github.com/meritlabs/achievement-engine/pkg/dto"
	"github.com/meritlabs/achievement-engine/pkg/services"
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

		if err != nil {
			c.Error(dto.BadRequestError{Message: err.Error()})
			return
		}

		if user == nil || user.Status != models.Approved {
			c.Error(dto.ForbiddenError{})
			return
		}

		token, err := s.CreateSession(*user)

		if err != nil {
			c.Error(dto.ForbiddenError{})
			return
		}

		c.JSON(http.StatusOK, dto.TokenSessionResponse{token})
	}
}
