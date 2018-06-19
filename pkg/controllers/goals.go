package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meritlabs/achievement-engine/pkg/db/models/goal"
)

func ListGoals(c *gin.Context) {
	c.JSON(http.StatusOK, goal.GetGoals())
}