package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ListAchievements(c *gin.Context) {

}

func GetAchievement(c *gin.Context) {
	slug := c.Param("slug")
	fmt.Println(slug)
}

func UpdateAchievement(c *gin.Context) {
	slug := c.Param("slug")
	fmt.Println(slug)
}
