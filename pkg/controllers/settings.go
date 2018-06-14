package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meritlabs/achievement-engine/pkg/db/models"
	"github.com/meritlabs/achievement-engine/pkg/db/stores"
	"github.com/meritlabs/achievement-engine/pkg/dto"
)

func GetSettings(store *stores.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*models.User)

		settings, err := store.GetUserSettings(user.ID)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, settings)
	}
}

func UpdateSettings(store *stores.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*models.User)

		var payload dto.SettingsPayload
		c.BindJSON(&payload)

		settings := models.Settings{
			UserID:                 user.ID,
			IsSetupTrackerEnabled:  payload.IsSetupTrackerEnabled,
			IsWelcomeDialogEnabled: payload.IsWelcomeDialogEnabled,
		}

		if err := store.UpdateUserSettings(user.ID, &settings); err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusAccepted, nil)
	}
}
