package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/meritlabs/achievement-engine/db"
	"github.com/meritlabs/achievement-engine/db/models"
)

func ListGoals(c *gin.Context) {
	session, err := db.WithDBSession()
	if err != nil {
		c.Error(err)
		return
	}
	defer session.Close()

	db := session.DB("achievement-engine").C("goals")

	var result []models.Goal

	err = db.Find(bson.M{}).All(&result)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
