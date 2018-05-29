package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meritlabs/achievement-engine/db/stores"
)

func ListGoals(store *stores.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		goals, err := store.ListGoals()
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, goals)
	}
}
