package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db/models"
	"github.com/meritlabs/achievement-engine/db/stores"
)

func ListAchievements(store *stores.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(models.User)

		achievements, err := store.GetAchievementsForUser(user.ID)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, achievements)
	}
}

func GetAchievement(store *stores.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(models.User)
		slug := c.Param("slug")

		if !bson.IsObjectIdHex(slug) {
			c.Error(fmt.Errorf("Slug %s is not valid object id", slug))
			return
		}

		achievement, err := store.GetAchievementForUser(user.ID, bson.ObjectIdHex(slug))
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, achievement)
	}
}

func UpdateAchievement(store *stores.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(models.User)
		fmt.Print("%v", user)
	}
}
